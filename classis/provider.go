package classis

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/josmo/terraform-provider-sampleclassis/classis/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Default:     "https://spot.classis.io",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLASSIS_URL", nil),
				Description: "url for classis",
			},
			"email": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLASSIS_EMAIL", nil),
				Description: "email for classis",
			},
			"password": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLASSIS_PASSWORD", nil),
				Description: "password for classis",
			},
		},
		ResourcesMap:  map[string]*schema.Resource{
			"sampleclassis_aws_spot_group": resourceAwsSpotGroupObject(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	return classis.NewClientWith(url, email, password)
}
