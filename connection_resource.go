package main

import "github.com/hashicorp/terraform/helper/schema"

func NewConnectionResource() *schema.Resource {
	return &schema.Resource{
		Exists:   connectionResourceExists,
		Create:   connectionResourceCreate,
		Read:     connectionResourceRead,
		Update:   connectionResourceUpdate,
		Delete:   connectionResourceDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   connectionResourceSchema(),
	}
}

func connectionResourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	_, _, err := metaConnectionService(meta).Read(d.Id(), nil)

	if err != nil && isApi404Err(err) {
		return false, nil
	} else {
		return err == nil, err
	}
}

func connectionResourceCreate(d *schema.ResourceData, meta interface{}) error {
	props := connectionResourceDataToProperties(d)

	if connection, _, err := metaConnectionService(meta).Create(props); err != nil {
		return err
	} else {
		d.SetId(connection.Id.(string))
		return nil
	}
}

func connectionResourceRead(d *schema.ResourceData, meta interface{}) error {
	if connection, _, err := metaConnectionService(meta).Read(d.Id(), nil); err != nil {
		return err
	} else {
		return connectionResourceDataFromConnection(d, connection)
	}
}

func connectionResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	props := connectionResourceDataToProperties(d)

	if _, _, err := metaConnectionService(meta).Update(d.Id(), props); err != nil {
		return err
	} else {
		return nil
	}
}

func connectionResourceDelete(d *schema.ResourceData, meta interface{}) error {
	if _, err := metaConnectionService(meta).Delete(d.Id()); err != nil {
		return err
	} else {
		return nil
	}
}
