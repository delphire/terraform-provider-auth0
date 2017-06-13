package main

import "github.com/hashicorp/terraform/helper/schema"

func NewClientResource() *schema.Resource {
	return &schema.Resource{
		Exists:   clientResourceExists,
		Create:   clientResourceCreate,
		Read:     clientResourceRead,
		Update:   clientResourceUpdate,
		Delete:   clientResourceDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   clientResourceSchema(),
	}
}

func clientResourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	_, _, err := metaClientService(meta).Read(d.Id(), nil)

	if err != nil && isApi404Err(err) {
		return false, nil
	} else {
		return err == nil, err
	}
}

func clientResourceCreate(d *schema.ResourceData, meta interface{}) error {
	props := clientResourceDataToProperties(d)

	if client, _, err := metaClientService(meta).Create(props); err != nil {
		return err
	} else {
		d.SetId(client.ClientId.(string))
		return nil
	}
}

func clientResourceRead(d *schema.ResourceData, meta interface{}) error {
	if client, _, err := metaClientService(meta).Read(d.Id(), nil); err != nil {
		return err
	} else {
		return clientResourceDataFromClient(d, client)
	}
}

func clientResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	props := clientResourceDataToProperties(d)

	if _, _, err := metaClientService(meta).Update(d.Id(), props); err != nil {
		return err
	} else {
		return nil
	}
}

func clientResourceDelete(d *schema.ResourceData, meta interface{}) error {
	if _, err := metaClientService(meta).Delete(d.Id()); err != nil {
		return err
	} else {
		return nil
	}
}
