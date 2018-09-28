package url

import (
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider for zip
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"url": dataURL(),
		},
	}
}

func dataURL() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Raw URL input",
			},
			"scheme": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL scheme (e.g. http)",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username of the URL (if given)",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password of the URL (if given)",
			},
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the host including the port",
			},
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the host excluding the port",
			},
			"port": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Port of the URL",
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path of the URL",
			},
			"query": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Query map",
			},
		},
		Read: func(d *schema.ResourceData, meta interface{}) error {
			rawurl := d.Get("url").(string)
			u, err := url.Parse(rawurl)
			if err != nil {
				return err
			}
			d.SetId(u.String())
			d.Set("host", u.Host)
			d.Set("hostname", u.Hostname())
			d.Set("port", u.Port())
			d.Set("path", u.Path)
			d.Set("scheme", u.Scheme)

			// auth
			d.Set("username", u.User.Username())
			pass, _ := u.User.Password()
			d.Set("password", pass)

			query := map[string]string{}
			values := u.Query()
			for key := range values {
				query[key] = values.Get(key)
			}
			d.Set("query", query)

			return nil
		},
	}
}
