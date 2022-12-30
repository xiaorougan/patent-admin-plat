package charts

import "encoding/json"

func init() {
	globalCharts[1001] = newChart1001()
}

type resp1001 struct {
	ReturnValue int `json:"ReturnValue"`
	Option      struct {
		Statistics [][]int  `json:"Statistics"`
		Classes    []string `json:"Classes"`
	} `json:"Option"`
}

type chart1001 struct {
	name string
}

func newChart1001() *chart1001 {
	return &chart1001{
		name: "省市申请量分析",
	}
}

func (c *chart1001) Serialize(params []byte) (string, error) {
	resp := &resp1001{}
	if err := json.Unmarshal(params, resp); err != nil {
		return "", err
	}

	data := make([]int, 0, len(resp.Option.Statistics))
	for _, s := range resp.Option.Statistics {
		data = append(data, s[0])
	}
	chart := genVerticalBarProfile(resp.Option.Classes, data, false)

	return chart, nil
}

func (c *chart1001) Name() string {
	return c.name
}
