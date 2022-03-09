// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouteCreate,
		Read:   resourceComputeRouteRead,
		Delete: resourceComputeRouteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRouteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dest_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The destination range of outgoing packets that this route applies to.
Only IPv4 is supported.`,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the
last character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The network that this route applies to.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property
when you create the resource.`,
			},
			"next_hop_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL to a gateway that should handle matching packets.
Currently, you can only specify the internet gateway, using a full or
partial valid URL:
* 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
* 'projects/project/global/gateways/default-internet-gateway'
* 'global/gateways/default-internet-gateway'
* The string 'default-internet-gateway'.`,
				ExactlyOneOf: []string{"next_hop_gateway", "next_hop_instance", "next_hop_ip", "next_hop_vpn_tunnel", "next_hop_ilb"},
			},
			"next_hop_ilb": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareIpAddressOrSelfLinkOrResourceName,
				Description: `The IP address or URL to a forwarding rule of type
loadBalancingScheme=INTERNAL that should handle matching
packets.

With the GA provider you can only specify the forwarding
rule as a partial or full URL. For example, the following
are all valid values:
* 10.128.0.56
* https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
* regions/region/forwardingRules/forwardingRule

When the beta provider, you can also specify the IP address
of a forwarding rule from the same VPC or any peered VPC.

Note that this can only be used when the destinationRange is
a public (non-RFC 1918) IP CIDR range.`,
				ExactlyOneOf: []string{"next_hop_gateway", "next_hop_instance", "next_hop_ip", "next_hop_vpn_tunnel", "next_hop_ilb"},
			},
			"next_hop_instance": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL to an instance that should handle matching packets.
You can specify this as a full or partial URL. For example:
* 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'
* 'projects/project/zones/zone/instances/instance'
* 'zones/zone/instances/instance'
* Just the instance name, with the zone in 'next_hop_instance_zone'.`,
				ExactlyOneOf: []string{"next_hop_gateway", "next_hop_instance", "next_hop_ip", "next_hop_vpn_tunnel", "next_hop_ilb"},
			},
			"next_hop_ip": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				Description:  `Network IP address of an instance that should handle matching packets.`,
				ExactlyOneOf: []string{"next_hop_gateway", "next_hop_instance", "next_hop_ip", "next_hop_vpn_tunnel", "next_hop_ilb"},
			},
			"next_hop_vpn_tunnel": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL to a VpnTunnel that should handle matching packets.`,
				ExactlyOneOf:     []string{"next_hop_gateway", "next_hop_instance", "next_hop_ip", "next_hop_vpn_tunnel", "next_hop_ilb"},
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The priority of this route. Priority is used to break ties in cases
where there is more than one matching route of equal prefix length.

In the case of two routes with equal prefix length, the one with the
lowest-numbered priority value wins.

Default value is 1000. Valid range is 0 through 65535.`,
				Default: 1000,
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: `A list of instance tags to which this route applies.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"next_hop_network": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `URL to a Network that should handle matching packets.`,
			},
			"next_hop_instance_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The zone of the instance specified in next_hop_instance. Omit if next_hop_instance is specified as a URL.",
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	destRangeProp, err := expandComputeRouteDestRange(d.Get("dest_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dest_range"); !isEmptyValue(reflect.ValueOf(destRangeProp)) && (ok || !reflect.DeepEqual(v, destRangeProp)) {
		obj["destRange"] = destRangeProp
	}
	descriptionProp, err := expandComputeRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRouteName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeRouteNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputeRoutePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("priority"); ok || !reflect.DeepEqual(v, priorityProp) {
		obj["priority"] = priorityProp
	}
	tagsProp, err := expandComputeRouteTags(d.Get("tags"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tags"); !isEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	nextHopGatewayProp, err := expandComputeRouteNextHopGateway(d.Get("next_hop_gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_gateway"); !isEmptyValue(reflect.ValueOf(nextHopGatewayProp)) && (ok || !reflect.DeepEqual(v, nextHopGatewayProp)) {
		obj["nextHopGateway"] = nextHopGatewayProp
	}
	nextHopInstanceProp, err := expandComputeRouteNextHopInstance(d.Get("next_hop_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_instance"); !isEmptyValue(reflect.ValueOf(nextHopInstanceProp)) && (ok || !reflect.DeepEqual(v, nextHopInstanceProp)) {
		obj["nextHopInstance"] = nextHopInstanceProp
	}
	nextHopIpProp, err := expandComputeRouteNextHopIp(d.Get("next_hop_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_ip"); !isEmptyValue(reflect.ValueOf(nextHopIpProp)) && (ok || !reflect.DeepEqual(v, nextHopIpProp)) {
		obj["nextHopIp"] = nextHopIpProp
	}
	nextHopVpnTunnelProp, err := expandComputeRouteNextHopVpnTunnel(d.Get("next_hop_vpn_tunnel"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_vpn_tunnel"); !isEmptyValue(reflect.ValueOf(nextHopVpnTunnelProp)) && (ok || !reflect.DeepEqual(v, nextHopVpnTunnelProp)) {
		obj["nextHopVpnTunnel"] = nextHopVpnTunnelProp
	}
	nextHopIlbProp, err := expandComputeRouteNextHopIlb(d.Get("next_hop_ilb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("next_hop_ilb"); !isEmptyValue(reflect.ValueOf(nextHopIlbProp)) && (ok || !reflect.DeepEqual(v, nextHopIlbProp)) {
		obj["nextHopIlb"] = nextHopIlbProp
	}

	lockName, err := replaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/peerings")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/routes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Route: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Route: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), isPeeringOperationInProgress)
	if err != nil {
		return fmt.Errorf("Error creating Route: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating Route", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Route: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Route %q: %#v", d.Id(), res)

	return resourceComputeRouteRead(d, meta)
}

func resourceComputeRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Route: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil, isPeeringOperationInProgress)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRoute %q", d.Id()))
	}

	res, err = resourceComputeRouteDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeRoute because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}

	if err := d.Set("dest_range", flattenComputeRouteDestRange(res["destRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("description", flattenComputeRouteDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("name", flattenComputeRouteName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("network", flattenComputeRouteNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("priority", flattenComputeRoutePriority(res["priority"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("tags", flattenComputeRouteTags(res["tags"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_gateway", flattenComputeRouteNextHopGateway(res["nextHopGateway"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_instance", flattenComputeRouteNextHopInstance(res["nextHopInstance"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_ip", flattenComputeRouteNextHopIp(res["nextHopIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_vpn_tunnel", flattenComputeRouteNextHopVpnTunnel(res["nextHopVpnTunnel"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_network", flattenComputeRouteNextHopNetwork(res["nextHopNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_ilb", flattenComputeRouteNextHopIlb(res["nextHopIlb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}

	return nil
}

func resourceComputeRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Route: %s", err)
	}
	billingProject = project

	lockName, err := replaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/peerings")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Route %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), isPeeringOperationInProgress)
	if err != nil {
		return handleNotFoundError(err, d, "Route")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting Route", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Route %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/routes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRouteDestRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouteDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouteName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouteNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRoutePriority(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeRouteTags(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeRouteNextHopGateway(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouteNextHopInstance(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRouteNextHopIp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouteNextHopVpnTunnel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRouteNextHopNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRouteNextHopIlb(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeRouteDestRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRoutePriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeRouteNextHopGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == "default-internet-gateway" {
		return replaceVars(d, config, "projects/{{project}}/global/gateways/default-internet-gateway")
	} else {
		return v, nil
	}
}

func expandComputeRouteNextHopInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == "" {
		return v, nil
	}
	val, err := parseZonalFieldValue("instances", v.(string), "project", "next_hop_instance_zone", d, config, true)
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return nil, err
	}

	nextInstance, err := config.NewComputeClient(userAgent).Instances.Get(val.Project, val.Zone, val.Name).Do()
	if err != nil {
		return nil, err
	}
	return nextInstance.SelfLink, nil
}

func expandComputeRouteNextHopIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNextHopVpnTunnel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("vpnTunnels", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for next_hop_vpn_tunnel: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRouteNextHopIlb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeRouteDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["nextHopInstance"]; ok {
		val, err := parseZonalFieldValue("instances", v.(string), "project", "next_hop_instance_zone", d, meta.(*Config), true)
		if err != nil {
			return nil, err
		}
		if err := d.Set("next_hop_instance_zone", val.Zone); err != nil {
			return nil, fmt.Errorf("Error setting next_hop_instance_zone: %s", err)
		}
		res["nextHopInstance"] = val.RelativeLink()
	}

	return res, nil
}