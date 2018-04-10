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

	err := genApiKey(key)
	if err != nil {
	   return err
	}

	d.SetId(key.Host)

	if err = d.Set("gen_access_key", key.GenAccessKey); err != nil {
	   return err
	}

	if err = d.Set("gen_secret_key", key.GenSecretKey); err != nil {
	   return err
	}

	return nil
}

func resourceApiKeyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceApiKeyUpdate(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceApiKeyDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
