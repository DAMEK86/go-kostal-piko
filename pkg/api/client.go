package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"encoding/json"
)

const (
	DefaultUsername string = "pvserver"
	DefaultPassword string = "pvwr"
)

/*Client abstracts http client*/
type Client interface {
	GetData(response *DxsRespone, dataPointsToRequest []int) error
	GetGeneralData(response *DxsRespone) error
	GetStatsData(response *DxsRespone) error
}

type client struct {
	httpClient        *http.Client
	username          string
	password          string
	url               string
	dataPoints        map[string]DataPoint
	generalDataPoints []int
	statsDataPoints   []int
}

/*NewClient creates a new instance of Client*/
func NewClient(httpClient *http.Client, baseURL, username, password string) Client {
	dataMapping := getDataPoints()
	return &client{
		httpClient:        httpClient,
		url:               fmt.Sprintf("http://%s/api/dxs.json", baseURL),
		username:          username,
		password:          password,
		dataPoints:        dataMapping,
		generalDataPoints: getGeneralDataPoints(dataMapping),
		statsDataPoints:   getStatsDataPoints(dataMapping),
	}
}

func (c *client) GetGeneralData(response *DxsRespone) error {
	return c.GetData(response, c.generalDataPoints)
}

func (c *client) GetStatsData(response *DxsRespone) error {
	return c.GetData(response, c.statsDataPoints)
}

func (c *client) GetData(response *DxsRespone, dataPointsToRequest []int) error {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	params := url.Values{}
	for i := 0; i < len(dataPointsToRequest); i++ {
		params.Add("dxsEntries", strconv.Itoa(dataPointsToRequest[i]))
	}
	req.URL.RawQuery = params.Encode()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)

	if err != nil {
		return fmt.Errorf("failed decode http response: %s", err)
	}

	for i := 0; i < len(response.Entries); i++ {
		dataPoint := c.getDataPointFromID(response.Entries[i].ID)
		response.Entries[i].FriendlyName = dataPoint.FriendlyName
		response.Entries[i].Unit = dataPoint.Unit
	}

	return nil
}

func (c *client) getDataPointFromID(id int) DataPoint {
	for _, v := range c.dataPoints {
		if v.ID == id {
			return v
		}
	}
	return DataPoint{}
}

func getDataPoints() map[string]DataPoint {
	return map[string]DataPoint{
		"pv_generator_dc_input_1_voltage": {
			ID:           33555202,
			FriendlyName: "pv_generator_dc_input_1_voltage",
			Unit:         "V",
		},
		"pv_generator_dc_input_1_current": {
			ID:           33555201,
			FriendlyName: "pv_generator_dc_input_1_current",
			Unit:         "A",
		},
		"pv_generator_dc_input_1_power": {
			ID:           33555203,
			FriendlyName: "pv_generator_dc_input_1_power",
			Unit:         "W",
		},
		"pv_generator_dc_input_2_voltage": {
			ID:           33555458,
			FriendlyName: "pv_generator_dc_input_2_voltage",
			Unit:         "W",
		},
		"pv_generator_dc_input_2_current": {
			ID:           33555457,
			FriendlyName: "pv_generator_dc_input_2_current",
			Unit:         "W",
		},
		"pv_generator_dc_input_2_power": {
			ID:           33555459,
			FriendlyName: "pv_generator_dc_input_2_power",
			Unit:         "W",
		},
		"house_home_consumption_covered_by_solar_generator": {
			ID:           83886336,
			FriendlyName: "house_home_consumption_covered_by_solar_generator",
			Unit:         "W",
		},
		"house_home_consumption_covered_by_battery": {
			ID:           83886592,
			FriendlyName: "house_home_consumption_covered_by_battery",
			Unit:         "W",
		},
		"house_home_consumption_covered_by_grid": {
			ID:           83886848,
			FriendlyName: "house_home_consumption_covered_by_grid",
			Unit:         "W",
		},
		"house_phase_selective_home_consumption_phase_1": {
			ID:           83887106,
			FriendlyName: "house_phase_selective_home_consumption_phase_1",
			Unit:         "W",
		},
		"house_phase_selective_home_consumption_phase_2": {
			ID:           83887362,
			FriendlyName: "house_phase_selective_home_consumption_phase_2",
			Unit:         "W",
		},
		"house_phase_selective_home_consumption_phase_3": {
			ID:           83887618,
			FriendlyName: "house_phase_selective_home_consumption_phase_3",
			Unit:         "W",
		},
		"grid_grid_parameters_output_power": {
			ID:           67109120,
			FriendlyName: "grid_grid_parameters_output_power",
			Unit:         "W",
		},
		"grid_grid_parameters_grid_frequency": {
			ID:           67110400,
			FriendlyName: "grid_grid_parameters_grid_frequency",
			Unit:         "Hz",
		},
		"grid_grid_parameters_cos": {
			ID:           67110656,
			FriendlyName: "grid_grid_parameters_cos",
			Unit:         "",
		},
		"grid_grid_parameters_limitation_on": {
			ID:           67110144,
			FriendlyName: "grid_grid_parameters_limitation_on",
			Unit:         "%",
		},
		"grid_phase_1_voltage": {
			ID:           67109378,
			FriendlyName: "grid_phase_1_voltage",
			Unit:         "V",
		},
		"grid_phase_1_current": {
			ID:           67109377,
			FriendlyName: "grid_phase_1_current",
			Unit:         "A",
		},
		"grid_phase_1_power": {
			ID:           67109379,
			FriendlyName: "grid_phase_1_power",
			Unit:         "W",
		},
		"grid_phase_2_voltage": {
			ID:           67109634,
			FriendlyName: "grid_phase_2_voltage",
			Unit:         "V",
		},
		"grid_phase_2_current": {
			ID:           67109633,
			FriendlyName: "grid_phase_2_current",
			Unit:         "A",
		},
		"grid_phase_2_power": {
			ID:           67109635,
			FriendlyName: "grid_phase_2_power",
			Unit:         "W",
		},
		"grid_phase_3_voltage": {
			ID:           67109890,
			FriendlyName: "grid_phase_3_voltage",
			Unit:         "V",
		},
		"grid_phase_3_current": {
			ID:           67109889,
			FriendlyName: "grid_phase_3_current",
			Unit:         "A",
		},
		"grid_phase_3_power": {
			ID:           67109891,
			FriendlyName: "grid_phase_3_power",
			Unit:         "W",
		},
		"stats_total_yield": {
			ID:           251658753,
			FriendlyName: "stats_total_yield",
			Unit:         "Wh",
		},
		"stats_total_operation_time": {
			ID:           251658496,
			FriendlyName: "stats_total_operation_time",
			Unit:         "h",
		},
		"stats_total_total_home_consumption": {
			ID:           251659009,
			FriendlyName: "stats_total_total_home_consumption",
			Unit:         "Wh",
		},
		"stats_total_self_consumption_kwh": {
			ID:           251659265,
			FriendlyName: "stats_total_self_consumption_kwh",
			Unit:         "Wh",
		},
		"stats_total_self_consumption_rate": {
			ID:           251659280,
			FriendlyName: "stats_total_self_consumption_rate",
			Unit:         "%",
		},
		"stats_total_degree_of_self_sufficiency": {
			ID:           251659281,
			FriendlyName: "stats_total_degree_of_self_sufficiency",
			Unit:         "%",
		},
		"stats_day_yield": {
			ID:           251658754,
			FriendlyName: "stats_day_yield",
			Unit:         "Wh",
		},
		"stats_day_total_home_consumption": {
			ID:           251659010,
			FriendlyName: "stats_day_total_home_consumption",
			Unit:         "Wh",
		},
		"stats_day_self_consumption_kwh": {
			ID:           251659266,
			FriendlyName: "stats_day_self_consumption_kwh",
			Unit:         "Wh",
		},
		"stats_day_self_consumption_rate": {
			ID:           251659278,
			FriendlyName: "stats_day_self_consumption_rate",
			Unit:         "%",
		},
		"stats_day_degree_of_self_sufficiency": {
			ID:           251659279,
			FriendlyName: "stats_day_degree_of_self_sufficiency",
			Unit:         "%",
		},
		"stats_operating_status": {
			ID: 16780032,
			FriendlyName: "stats_operating_status",
			Unit: "",
		},
	}
}

func getGeneralDataPoints(data map[string]DataPoint) []int {
	dataPoints := make([]int, 0)
	dataPoints = append(dataPoints, data["pv_generator_dc_input_1_voltage"].ID)
	dataPoints = append(dataPoints, data["pv_generator_dc_input_1_current"].ID)
	dataPoints = append(dataPoints, data["pv_generator_dc_input_1_power"].ID)
	dataPoints = append(dataPoints, data["pv_generator_dc_input_2_voltage"].ID)
	dataPoints = append(dataPoints, data["pv_generator_dc_input_2_current"].ID)
	dataPoints = append(dataPoints, data["pv_generator_dc_input_2_power"].ID)
	dataPoints = append(dataPoints, data["house_home_consumption_covered_by_solar_generator"].ID)
	dataPoints = append(dataPoints, data["house_home_consumption_covered_by_grid"].ID)
	dataPoints = append(dataPoints, data["house_phase_selective_home_consumption_phase_1"].ID)
	dataPoints = append(dataPoints, data["house_phase_selective_home_consumption_phase_2"].ID)
	dataPoints = append(dataPoints, data["house_phase_selective_home_consumption_phase_3"].ID)
	dataPoints = append(dataPoints, data["grid_grid_parameters_output_power"].ID)
	dataPoints = append(dataPoints, data["grid_grid_parameters_grid_frequency"].ID)
	dataPoints = append(dataPoints, data["grid_grid_parameters_cos"].ID)
	dataPoints = append(dataPoints, data["grid_phase_1_voltage"].ID)
	dataPoints = append(dataPoints, data["grid_phase_1_current"].ID)
	dataPoints = append(dataPoints, data["grid_phase_1_power"].ID)
	dataPoints = append(dataPoints, data["grid_phase_2_voltage"].ID)
	dataPoints = append(dataPoints, data["grid_phase_2_current"].ID)
	dataPoints = append(dataPoints, data["grid_phase_2_power"].ID)
	dataPoints = append(dataPoints, data["grid_phase_3_voltage"].ID)
	dataPoints = append(dataPoints, data["grid_phase_3_current"].ID)
	dataPoints = append(dataPoints, data["grid_phase_3_power"].ID)
	return dataPoints
}

func getStatsDataPoints(data map[string]DataPoint) []int {
	dataPoints := make([]int, 0)
	dataPoints = append(dataPoints, data["stats_total_yield"].ID)
	dataPoints = append(dataPoints, data["stats_total_operation_time"].ID)
	dataPoints = append(dataPoints, data["stats_total_total_home_consumption"].ID)
	dataPoints = append(dataPoints, data["stats_total_self_consumption_kwh"].ID)
	dataPoints = append(dataPoints, data["stats_total_self_consumption_rate"].ID)
	dataPoints = append(dataPoints, data["stats_total_degree_of_self_sufficiency"].ID)
	dataPoints = append(dataPoints, data["stats_day_yield"].ID)
	dataPoints = append(dataPoints, data["stats_day_total_home_consumption"].ID)
	dataPoints = append(dataPoints, data["stats_day_self_consumption_kwh"].ID)
	dataPoints = append(dataPoints, data["stats_day_self_consumption_rate"].ID)
	dataPoints = append(dataPoints, data["stats_day_degree_of_self_sufficiency"].ID)
	dataPoints = append(dataPoints, data["stats_operating_status"].ID)

	return dataPoints
}
