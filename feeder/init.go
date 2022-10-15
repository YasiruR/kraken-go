package feeder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/YasiruR/feeder/requests"
	"github.com/YasiruR/feeder/responses"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

const (
	assetInfoEndpoint = `https://api.kraken.com/0/public/Assets`
)

type Feeder struct {
	api    string
	client *http.Client
	*websocket.Dialer
}

func NewFeeder(api string) *Feeder {
	return &Feeder{api: api, client: &http.Client{}}
}

func (f *Feeder) AssetInfo() error {
	res, err := f.client.Get(assetInfoEndpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var assets responses.GetAssetInfo
	err = json.Unmarshal(data, &assets)
	if err != nil {
		return err
	}

	return nil
}

func (f *Feeder) Subscribe(pair string) error {
	conn, _, err := websocket.DefaultDialer.Dial(`wss://ws.kraken.com`, http.Header{})
	if err != nil {
		return err
	}
	defer conn.Close()

	var connRes responses.Connection
	err = conn.ReadJSON(&connRes)
	if err != nil {
		return err
	}

	if connRes.Event != `systemStatus` || connRes.Status != `online` {
		return fmt.Errorf(`init connection failed (event: %s, status: %s)`, connRes.Event, connRes.Status)
	}

	var req requests.Subscribe
	req.Event = `subscribe`
	req.Pair = []string{pair}
	req.Subscription.Name = `ticker`

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err = conn.WriteMessage(1, data); err != nil {
		return err
	}

	if err = conn.ReadJSON(&connRes); err != nil {
		return err
	}

	if connRes.Event != `subscriptionStatus` || connRes.Pair != pair || connRes.Status != `subscribed` {
		return fmt.Errorf(`subscription res invalid (event: %s, pair: %s, status: %s)`, connRes.Event, connRes.Pair, connRes.Status)
	}

	var res []interface{}
	for {
		_, data, err = conn.ReadMessage()
		if err != nil {
			return err
		}

		if bytes.Compare([]byte("{\"event\":\"heartbeat\"}"), data) == 0 {
			continue
		}

		if err = json.Unmarshal(data, &res); err != nil {
			return err
		}

		fmt.Println(res)
	}

}

type level struct {
	price  string
	volume string
}
