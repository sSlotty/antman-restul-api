package services

import (
	"antman/restful-api/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"net/http"
)

func GetCircuit(c *gin.Context) {
	id := c.Param("id")

	var antmanModel models.AntmanModel
	client := req.C()
	resp, err := client.R().Get("http://localhost:9200/antman-index-*/_search?size=1000&q=circuit:" + id)

	body := resp.String()
	err = json.Unmarshal([]byte(body), &antmanModel)
	if err != nil {
		c.JSON(500, gin.H{
			"data": err.Error(),
		})
		return
	}
	painscore := 0
	lowwifitxrx := 0
	disconnect := 0
	t3_wifi5g_issue := 0
	verycloseap := 0
	multiplenetwork := 0
	lowlanspeed := 0
	lanerror := 0
	internetdisconnect := 0
	fttxpacketdrop := 0
	lowfttxpower := 0
	hightemp := 0
	highmem := 0
	highcpu := 0
	reboot := 0
	adaptor_fault_rate := 0
	fiberflapping := 0
	powerloss := 0

	for i := 0; i < len(antmanModel.Hits.Hits); i++ {
		painscore += antmanModel.Hits.Hits[i].Source.Painscore
		lowwifitxrx += antmanModel.Hits.Hits[i].Source.Lowwifitxrx
		disconnect += antmanModel.Hits.Hits[i].Source.Disconnect
		t3_wifi5g_issue += antmanModel.Hits.Hits[i].Source.T3Wifi5GIssue
		verycloseap += antmanModel.Hits.Hits[i].Source.Verycloseap
		multiplenetwork += antmanModel.Hits.Hits[i].Source.Multiplenetwork
		lowlanspeed += antmanModel.Hits.Hits[i].Source.Lowlanspeed
		lanerror += antmanModel.Hits.Hits[i].Source.Lanerror
		internetdisconnect += antmanModel.Hits.Hits[i].Source.Internetdisconnect
		fttxpacketdrop += antmanModel.Hits.Hits[i].Source.Fttxpacketdrop
		lowfttxpower += antmanModel.Hits.Hits[i].Source.Lowfttxpower
		hightemp += antmanModel.Hits.Hits[i].Source.Hightemp
		highmem += antmanModel.Hits.Hits[i].Source.Highmem
		highcpu += antmanModel.Hits.Hits[i].Source.Highcpu
		reboot += antmanModel.Hits.Hits[i].Source.Reboot
		adaptor_fault_rate += antmanModel.Hits.Hits[i].Source.AdaptorFaultRate
		fiberflapping += antmanModel.Hits.Hits[i].Source.Fiberflapping
		powerloss += antmanModel.Hits.Hits[i].Source.Powerloss
	}

	data := map[string]int{
		"data_total":         len(antmanModel.Hits.Hits),
		"painscore":          painscore,
		"lowwifitxrx":        lowwifitxrx,
		"disconnect":         disconnect,
		"t3_wifi5g_issue":    t3_wifi5g_issue,
		"verycloseap":        verycloseap,
		"multiplenetwork":    multiplenetwork,
		"lowlanspeed":        lowlanspeed,
		"lanerror":           lanerror,
		"internetdisconnect": internetdisconnect,
		"fttxpacketdrop":     fttxpacketdrop,
		"lowfttxpower":       lowfttxpower,
		"hightemp":           hightemp,
		"highmem":            highmem,
		"highcpu":            highcpu,
		"reboot":             reboot,
		"adaptor_fault_rate": adaptor_fault_rate,
		"fiberflapping":      fiberflapping,
		"powerloss":          powerloss,
	}

	c.JSON(
		200,
		gin.H{"status_code": http.StatusOK, "data": data},
	)
}
