package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceApiKey() *schema.Resource {
        return &schema.Resource{
                Create: resourceApiKeyCreate,
                Read:   resourceApiKeyRead,
                Update: resourceApiKeyUpdate,
                Delete: resourceApiKeyDelete,

                Schema: map[string]*schema.Schema{
                        "host": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:	  schema.TypeString,
				Optional: true,
			},
			"gen_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gen_secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id":	  &schema.Schema{
				Type:	  schema.TypeString,
				Optional: true,
				Default:  "1a1",
			},
                },
	}
}

func resourceApiKeyCreate(d *schema.ResourceData, m interface{}) error {
	key := &ApiKeyDescriptor{}

	key.Name = d.Get("name").(string)
	key.Host = d.Get("host").(string)

	if _akey, ok := d.GetOk("access_key"); ok {
		key.AccessKey = _akey.(string)
	}

	if _skey, ok := d.GetOk("secret_key"); ok {
		key.SecretKey = _skey.(string)
	}

	if _desc, ok := d.GetOk("description"); ok {
		key.Description = _desc.(string)
	}

	if _proj, ok := d.GetOk("project_id"); ok {
		key.ProjectId = _proj.(string)
	}

	err := genApiKey(key)
	if err != nil {
		return err
	}

	if err := d.Set("gen_access_key", key.GenAccessKey); err != nil {
		return err
	}

	if err := d.Set("gen_secret_key", key.GenSecretKey); err != nil {
		return err
	}

	d.SetId(key.UUID)

	return nil
}

func resourceApiKeyUpdate(d *schema.ResourceData, m interface{}) error {
	key := &ApiKeyDescriptor{
	        UUID:           d.Id(),
	        Host:           d.Get("host").(string),
	        GenAccessKey:   d.Get("gen_access_key").(string),
	        GenSecretKey:   d.Get("gen_secret_key").(string),
	        ProjectId:      d.Get("project_id").(string),
		Name:		d.Get("name").(string),
		Description:	d.Get("description").(string),
        }
	d.Partial(true)

	if d.HasChange("name") {
		_, _dName := d.GetChange("name")
		key.Name = _dName.(string)
		if err := updateApiKey(key); err != nil {
		   return err
		}

		d.SetPartial("name")
	}

	if d.HasChange("description") {
		_, _dDesc := d.GetChange("description")
		key.Description = _dDesc.(string)
		if err := updateApiKey(key); err != nil {
		   return err
		}

		d.SetPartial("description")
	}
	d.Partial(false)

	return nil
}

func resourceApiKeyRead(d *schema.ResourceData, m interface{}) error {
	key := &ApiKeyDescriptor{
		UUID:		d.Id(),
		Host:		d.Get("host").(string),
		ProjectId:	d.Get("project_id").(string),
		GenAccessKey:   d.Get("gen_access_key").(string),
                GenSecretKey:   d.Get("gen_secret_key").(string),
	}

	if err := readApiKey(key); err != nil {
		d.SetId("")
		return nil
	}

	err := d.Set("name", key.Name)
	if err != nil {
		return err
	}

	err = d.Set("description", key.Description)
	if err != nil {
		return err
	}

        return nil
}

func resourceApiKeyDelete(d *schema.ResourceData, m interface{}) error {
	key := &ApiKeyDescriptor{
		UUID:		d.Id(),
		Host:		d.Get("host").(string),
		GenAccessKey:	d.Get("gen_access_key").(string),
		GenSecretKey:	d.Get("gen_secret_key").(string),
		ProjectId:	d.Get("project_id").(string),
	}
	err := delApiKey(key)

	if err != nil {
	   return err
	}

	d.SetId("")
        return nil
}
