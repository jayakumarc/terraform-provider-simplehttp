package simplehttp

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
	"time"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"body": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	url := d.Get("url").(string)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error while making request to: %s. Error: %s", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request to %s is not successful. Response status code: %d", url, resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading the response body. Error: %s", err)
	}

	d.Set("body", string(bytes))
	d.SetId(time.Now().UTC().String())
	return nil
}
