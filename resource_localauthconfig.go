package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceLocalAuthConfig() *schema.Resource {
        return &schema.Resource{
                Create: resourceLocalAuthConfigCreate,
                Read:   resourceLocalAuthConfigRead,
                Update: resourceLocalAuthConfigUpdate,
                Delete: resourceLocalAuthConfigDelete,

                Schema: map[string]*schema.Schema{
                        "host": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id":     &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1a1",
			},
			"realname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:	  schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"accessmode":	  &schema.Schema{
				Type:	  schema.TypeString,
				Optional: true,
				Default:  "unrestricted",
			},
                },
	}
}

func resourceLocalAuthConfigCreate(d *schema.ResourceData, m interface{}) error {
	key := &LocalAuthConfigDescriptor{}

	key.Host = d.Get("host").(string)
	key.AccessKey = d.Get("access_key").(string)
	key.SecretKey = d.Get("secret_key").(string)

	if _proj, ok := d.GetOk("project_id"); ok {
		key.ProjectId = _proj.(string)
	}

	key.Username = d.Get("username").(string)
	key.Password = d.Get("password").(string)

	if _rname, ok := d.GetOk("realname"); ok {
		key.Realname = _rname.(string)
	}

	key.Enabled = d.Get("enabled").(bool)

	if _accessmode, ok := d.GetOk("accessmode"); ok {
		key.Accessmode = _accessmode.(string)
	}

	err := genLocalAuthConfig(key)
	if err != nil {
	   return err
	}

	d.SetId(key.UUID)

	return nil
}

func resourceLocalAuthConfigUpdate(d *schema.ResourceData, m interface{}) error {

   key := &LocalAuthConfigDescriptor{
		UUID:		d.Id(),
		Host:		d.Get("host").(string),
		AccessKey:	d.Get("access_key").(string),
		SecretKey:	d.Get("secret_key").(string),
		Username:	d.Get("username").(string),
		Password:       d.Get("password").(string),
		Realname:	d.Get("realname").(string),
		Enabled:	d.Get("enabled").(bool),
		Accessmode:	d.Get("accessmode").(string),

	}

	d.Partial(true)
	toChange := false

	if d.HasChange("username") {
		_, _dUname := d.GetChange("username")
		key.Username = _dUname.(string)
		toChange = true
	}

	if d.HasChange("password") {
		_, _dPname := d.GetChange("password")
		key.Password = _dPname.(string)
		toChange = true
	}

	if d.HasChange("realname") {
		_, _dRname := d.GetChange("realname")
		key.Realname = _dRname.(string)
		toChange = true
	}

	if d.HasChange("enabled") {
		_, _dEnable  := d.GetChange("enabled")
		key.Enabled = _dEnable.(bool)
		toChange = true
	}

	if d.HasChange("accessmode") {
		_, _dAccess :=  d.GetChange("accessmode")
		key.Accessmode = _dAccess.(string)
		toChange = true
	}

	if toChange == true {
	        err := genLocalAuthConfig(key)
	        if err != nil {
			return err
	       }
	}

	d.SetPartial("username")
	d.SetPartial("password")
	d.SetPartial("enabled")
	d.SetPartial("realname")
	d.SetPartial("accessmode")

	d.Partial(false)

//        d.SetId(key.UUID)

		return nil
}

func resourceLocalAuthConfigRead(d *schema.ResourceData, m interface{}) error {
/*	key := &LocalAuthConfigDescriptor{
		UUID:		d.Id(),
		Host:		d.Get("host").(string),
		ProjectId:	d.Get("project_id").(string),
		GenAccessKey:   d.Get("gen_access_key").(string),
                GenSecretKey:   d.Get("gen_secret_key").(string),
	}

	if err := readApiKey(key); err != nil {
		return err
	}

	err := d.Set("name", key.Name)
	if err != nil {
		d.SetId("")
		return nil
//		return err
	}

	err = d.Set("description", key.Description)
	if err != nil {
		return err
	}
*/
        return nil
}

func resourceLocalAuthConfigDelete(d *schema.ResourceData, m interface{}) error {
	key := &LocalAuthConfigDescriptor{
		UUID:           d.Id(),
		Host:           d.Get("host").(string),
		AccessKey:      d.Get("access_key").(string),
		SecretKey:      d.Get("secret_key").(string),
		Username:       "admin",
		Password:       "admin",
		Realname:       "",
		Enabled:        false,
		Accessmode:     "unrestricted",
	}


	err := genLocalAuthConfig(key)

	if err != nil {
	   return err
	}

	d.SetId("")
        return nil
}
