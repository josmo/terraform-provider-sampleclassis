package classis

import (
	"github.com/josmo/terraform-provider-sampleclassis/classis/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsSpotGroupObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsSpotGroupObjectCreate,
		Read:   resourceAwsSpotGroupObjectRead,
		Delete: resourceAwsSpotGroupObjectDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iam_fleet_role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"desired_qty": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"quantity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_device_size": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_types": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_groups": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceAwsSpotGroupObjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*classis.Client)
	var spotGroup = classis.SpotGroup{}
	spotGroup.Name = d.Get("group_name").(string)
	spotGroup.Region = d.Get("region").(string)
	spotGroup.DesiredQty = d.Get("desired_qty").(string)
	spotGroup.Quantity = d.Get("quantity").(string)
	spotGroup.IamFleetRole = d.Get("iam_fleet_role").(string)
	spotGroup.Vpc = d.Get("vpc_id").(string)
	if v, ok := d.GetOk("instance_types"); ok {
		instances := make([]string, len(v.([]interface{})))
		for _, element := range v.([]interface{}) {
			instances = append(instances, element.(string))
		}
		spotGroup.TypesSelected = instances
	}
	var launchSpecification = classis.LaunchSpecification{}
	launchSpecification.ImageId = d.Get("image_id").(string)
	launchSpecification.SubnetId = d.Get("subnet_id").(string)
	launchSpecification.KeyName = d.Get("key_name").(string)
	launchSpecification.DefaultDeviceSize = d.Get("default_device_size").(string)
	if v, ok := d.GetOk("security_groups"); ok {
		var sendSecurityGroups = []classis.SecurityGroup{}
		for _, element := range v.([]interface{}) {
			var securityGroup = classis.SecurityGroup{ element.(string)}
			sendSecurityGroups = append(sendSecurityGroups, securityGroup)
		}
		launchSpecification.SecurityGroups = sendSecurityGroups
	}

	spotGroup.LaunchSpecification = launchSpecification


	id, err := client.CreateSpotGroup(spotGroup)
	d.SetId(id)
	if err != nil {
		return err
	}
	return nil
}

func resourceAwsSpotGroupObjectRead(d *schema.ResourceData, meta interface{}) error {
    //TODO: This really should hit a get API to see what the current state is.  Something like the following sudo code
    // client := meta.(*classis.Client)
    // spotGroup, err := client.ReadSpotGroup(d.Id())
    // if err then d.SetId("")
    // else set and changed values
	return nil
}

func resourceAwsSpotGroupObjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*classis.Client)
	return client.DeleteSpotGroup(d.Id())
}
