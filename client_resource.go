package main

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
)

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
	LogPrintf("[%s] ResourceExists", "INFO")
	_, _, err := metaClientService(meta).Read(d.Id(), nil)

	if err != nil && isApi404Err(err) {
		return false, nil
	} else {
		return err == nil, err
	}
}

func clientResourceCreate(d *schema.ResourceData, meta interface{}) error {
	LogPrintf("[%s] ResourceCreate", "INFO")
	props := clientResourceDataToProperties(d)

	bytes, _ := json.MarshalIndent(props, "", "	")
	LogPrintf("[%s] Create: Properties JSON\n %s", "INFO", string(bytes))

	if client, _, err := metaClientService(meta).Create(props); err != nil {
		return err
	} else {
		bytes, _ := json.MarshalIndent(client, "", "	")
		LogPrintf("[%s] Update: Client JSON\n %s", "INFO", string(bytes))
		d.SetId(client.ClientId.(string))
		return nil
		// return clientResourceDataFromClient(d, client)
	}
}

func clientResourceRead(d *schema.ResourceData, meta interface{}) error {
	LogPrintf("[%s] ResourceRead", "INFO")
	if client, _, err := metaClientService(meta).Read(d.Id(), nil); err != nil {
		return err
	} else {
		bytes, _ := json.MarshalIndent(client, "", "	")
		LogPrintf("[%s] Read: Client JSON\n %s", "INFO", string(bytes))
		return clientResourceDataFromClient(d, client)
	}
}

func clientResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	LogPrintf("[%s] ResourceUpdate", "INFO")
	props := clientResourceDataToProperties(d)

	bytes, _ := json.MarshalIndent(props, "", "	")
	LogPrintf("[%s] Update: Properties JSON\n %s", "INFO", string(bytes))

	if client, _, err := metaClientService(meta).Update(d.Id(), props); err != nil {
		return err
	} else {
		bytes, _ := json.MarshalIndent(client, "", "	")
		LogPrintf("[%s] Update: Client JSON\n %s", "INFO", string(bytes))
		return nil
		// return clientResourceDataFromClient(d, client)
	}
}

func clientResourceDelete(d *schema.ResourceData, meta interface{}) error {
	LogPrintf("[%s] ResourceDelete", "INFO")
	if _, err := metaClientService(meta).Delete(d.Id()); err != nil {
		return err
	} else {
		return nil
	}
}
