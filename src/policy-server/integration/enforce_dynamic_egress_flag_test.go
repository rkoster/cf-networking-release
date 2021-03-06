package integration_test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"policy-server/config"
	"policy-server/integration/helpers"
	"policy-server/psclient"
	"strings"

	"code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport/metrics"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport/ports"
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnforceExperimentalDynamicEgressPolicies Flag", func() {
	var (
		sessions                  []*gexec.Session
		conf                      config.Config
		internalConf              config.InternalConfig
		dbConf                    db.Config
		tlsConfig                 *tls.Config
		policyServerConfs         []config.Config
		policyServerInternalConfs []config.InternalConfig
		fakeMetron                metrics.FakeMetron
	)

	BeforeEach(func() {
		fakeMetron = metrics.NewFakeMetron()
		dbConf = testsupport.GetDBConfig()
		dbConf.Timeout = 5
		dbConf.DatabaseName = fmt.Sprintf("internal_api_test_node_%d", ports.PickAPort())

		cert, err := tls.LoadX509KeyPair("fixtures/client.crt", "fixtures/client.key")
		Expect(err).NotTo(HaveOccurred())

		clientCACert, err := ioutil.ReadFile("fixtures/netman-ca.crt")
		Expect(err).NotTo(HaveOccurred())

		clientCertPool := x509.NewCertPool()
		clientCertPool.AppendCertsFromPEM(clientCACert)

		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      clientCertPool,
		}
		tlsConfig.BuildNameToCertificate()

		template, internalTemplate := helpers.DefaultTestConfig(dbConf, fakeMetron.Address(), "fixtures")
		template.TagLength = 2
		internalTemplate.TagLength = 2
		internalTemplate.EnforceExperimentalDynamicEgressPolicies = false
		policyServerConfs = configurePolicyServers(template, 1)
		policyServerInternalConfs = configureInternalPolicyServers(internalTemplate, 1)
		sessions = startPolicyAndInternalServers(policyServerConfs, policyServerInternalConfs)
		conf = policyServerConfs[0]
		internalConf = policyServerInternalConfs[0]

		body := strings.NewReader(`{
			"policies": [
				{"source": { "id": "app1" }, "destination": { "id": "app2", "protocol": "tcp", "ports": { "start": 8080, "end": 8080 } } },
				{"source": { "id": "app3" }, "destination": { "id": "app1", "protocol": "tcp", "ports": { "start": 9999, "end": 9999 } } },
				{"source": { "id": "app3" }, "destination": { "id": "app2", "protocol": "tcp", "ports": { "start": 3333, "end": 4444 } } },
				{"source": { "id": "app3" }, "destination": { "id": "app4", "protocol": "tcp", "ports": { "start": 3333, "end": 3333 } } }
			]
		}`)

		helpers.MakeAndDoRequest(
			"POST",
			fmt.Sprintf("http://%s:%d/networking/v1/external/policies", conf.ListenHost, conf.ListenPort),
			nil,
			body,
		)

		logger := lagertest.NewTestLogger("internal_api_test")
		client := psclient.NewClient(logger, http.DefaultClient, fmt.Sprintf("http://%s:%d", conf.ListenHost, conf.ListenPort))

		createdDestinations, err := client.CreateDestinations("valid-token", psclient.Destination{
			IPs:         []psclient.IPRange{{Start: "10.27.1.1", End: "10.27.1.2"}},
			Ports:       []psclient.Port{{Start: 8080, End: 8081}},
			Protocol:    "tcp",
			Name:        "dest-1",
			Description: "dest-1-desc",
		}, psclient.Destination{
			IPs:         []psclient.IPRange{{Start: "10.27.1.3", End: "10.27.1.3"}},
			Ports:       []psclient.Port{{Start: 8080, End: 8081}},
			Protocol:    "tcp",
			Name:        "dest-2",
			Description: "dest-2-desc",
		})
		Expect(err).ToNot(HaveOccurred())

		_, err = client.CreateEgressPolicy(psclient.EgressPolicy{
			Source:      psclient.EgressPolicySource{ID: "live-app-1-guid"},
			Destination: psclient.Destination{GUID: createdDestinations[0].GUID},
		}, "valid-token")
		Expect(err).ToNot(HaveOccurred())
		_, err = client.CreateEgressPolicy(psclient.EgressPolicy{
			Source:      psclient.EgressPolicySource{ID: "live-space-1-guid", Type: "space"},
			Destination: psclient.Destination{GUID: createdDestinations[1].GUID},
		}, "valid-token")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		stopPolicyServers(sessions, policyServerConfs)

		Expect(fakeMetron.Close()).To(Succeed())
	})

	It("prevents the internal api from returning the egress policies", func() {
		resp := helpers.MakeAndDoHTTPSRequest(
			"GET",
			fmt.Sprintf("https://%s:%d/networking/v1/internal/policies?id=app1,app2,live-space-1-guid,live-app-1-guid", internalConf.ListenHost, internalConf.InternalListenPort),
			nil,
			tlsConfig,
		)
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		responseBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(responseBytes)).To(WithTransform(replaceGUID, MatchJSON(`{"total_policies": 3,
			"policies": [
				{"source": { "id": "app1", "tag": "0001" }, "destination": { "id": "app2", "tag": "0002", "protocol": "tcp", "ports": {"start": 8080, "end": 8080 } } },
				{"source": { "id": "app3", "tag": "0003" }, "destination": { "id": "app1", "tag": "0001", "protocol": "tcp", "ports": {"start": 9999, "end": 9999 } } },
				{"source": { "id": "app3", "tag": "0003" }, "destination": { "id": "app2", "tag": "0002", "protocol": "tcp", "ports": { "start": 3333, "end": 4444 } } }
			]
		}`)))
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Eventually(fakeMetron.AllEvents, "5s").Should(ContainElement(
			HaveName("InternalPoliciesRequestTime"),
		))
		Eventually(fakeMetron.AllEvents, "5s").Should(ContainElement(
			HaveName("StoreAllSuccessTime"),
		))
	})
})
