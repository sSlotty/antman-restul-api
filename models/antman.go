package models

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Response struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

type AntmanModel struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Date                        string  `json:"date"`
				AdaptorFaultRate            int     `json:"adaptor_fault_rate"`
				Circuit                     string  `json:"circuit"`
				Datatotal                   int     `json:"datatotal"`
				Disconnect                  int     `json:"disconnect"`
				Fiberflapping               int     `json:"fiberflapping"`
				Fttxpacketdrop              int     `json:"fttxpacketdrop"`
				Highcpu                     int     `json:"highcpu"`
				Highmem                     int     `json:"highmem"`
				Hightemp                    int     `json:"hightemp"`
				InfoAverageCPU              float64 `json:"info_average_cpu"`
				InfoAverageDevicesCount2GQ1 int     `json:"info_average_devices_count2g_q1"`
				InfoAverageDevicesCount2GQ2 int     `json:"info_average_devices_count2g_q2"`
				InfoAverageDevicesCount2GQ3 int     `json:"info_average_devices_count2g_q3"`
				InfoAverageDevicesCount2GQ4 int     `json:"info_average_devices_count2g_q4"`
				InfoAverageDevicesCount5GQ1 int     `json:"info_average_devices_count5g_q1"`
				InfoAverageDevicesCount5GQ2 int     `json:"info_average_devices_count5g_q2"`
				InfoAverageDevicesCount5GQ3 int     `json:"info_average_devices_count5g_q3"`
				InfoAverageDevicesCount5GQ4 int     `json:"info_average_devices_count5g_q4"`
				InfoAverageDevicesReconnect float64 `json:"info_average_devices_reconnect"`
				InfoAverageFttxpower        float64 `json:"info_average_fttxpower"`
				InfoAverageMem              float64 `json:"info_average_mem"`
				InfoAverageRssi2G           float64 `json:"info_average_rssi_2g"`
				InfoAverageRssi5G           int     `json:"info_average_rssi_5g"`
				InfoAverageSnr2G            float64 `json:"info_average_snr_2g"`
				InfoAverageSnr5G            int     `json:"info_average_snr_5g"`
				InfoAverageTemp             float64 `json:"info_average_temp"`
				InfoAverageTxrxRate2G       int     `json:"info_average_txrx_rate_2g"`
				InfoAverageTxrxRate5G       int     `json:"info_average_txrx_rate_5g"`
				InfoBandsteering            int     `json:"info_bandsteering"`
				InfoMaxFttxpower            int     `json:"info_max_fttxpower"`
				InfoMinFttxpower            float64 `json:"info_min_fttxpower"`
				Internetdisconnect          int     `json:"internetdisconnect"`
				Lanerror                    int     `json:"lanerror"`
				Lowfttxpower                int     `json:"lowfttxpower"`
				Lowlanspeed                 int     `json:"lowlanspeed"`
				Lowwifitxrx                 int     `json:"lowwifitxrx"`
				Model                       string  `json:"model"`
				Multiplenetwork             int     `json:"multiplenetwork"`
				No5Ghz                      int     `json:"no_5_ghz"`
				Painscore                   int     `json:"painscore"`
				Powerloss                   int     `json:"powerloss"`
				Reboot                      int     `json:"reboot"`
				Serial                      string  `json:"serial"`
				T3Wifi5GIssue               int     `json:"t_3_wifi_5_g_issue"`
				Toomanydevices              int     `json:"toomanydevices"`
				Toomanyinterference2G       int     `json:"toomanyinterference_2_g"`
				Toomanyinterference5G       int     `json:"toomanyinterference_5_g"`
				Version                     string  `json:"version"`
				Verycloseap                 int     `json:"verycloseap"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
