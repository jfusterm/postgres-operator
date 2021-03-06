package kubeapi

/*
 Copyright 2017-2018 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	log "github.com/Sirupsen/logrus"
	crv1 "github.com/crunchydata/postgres-operator/apis/cr/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
)

// GetpgupgradesBySelector gets a list of pgupgrades by selector
func GetpgupgradesBySelector(client *rest.RESTClient, upgradeList *crv1.PgupgradeList, selector, namespace string) error {

	var err error

	myselector := labels.Everything()

	if selector != "" {
		myselector, err = labels.Parse(selector)
		if err != nil {
			log.Error("could not parse selector value ")
			log.Error(err)
			return err
		}
	}

	log.Debug("myselector is " + myselector.String())

	err = client.Get().
		Resource(crv1.PgupgradeResourcePlural).
		Namespace(namespace).
		Param("labelSelector", myselector.String()).
		Do().
		Into(upgradeList)
	if err != nil {
		log.Error("error getting list of upgrades " + err.Error())
	}

	return err
}

// Getpgupgrades gets a list of pgupgrades
func Getpgupgrades(client *rest.RESTClient, upgradeList *crv1.PgupgradeList, namespace string) error {

	err := client.Get().
		Resource(crv1.PgupgradeResourcePlural).
		Namespace(namespace).
		Do().Into(upgradeList)
	if err != nil {
		log.Error("error getting list of upgrades " + err.Error())
		return err
	}

	return err
}

// Getpgupgrade gets a pgupgrade by name
func Getpgupgrade(client *rest.RESTClient, upgrade *crv1.Pgupgrade, name, namespace string) (bool, error) {

	err := client.Get().
		Resource(crv1.PgupgradeResourcePlural).
		Namespace(namespace).
		Name(name).
		Do().Into(upgrade)
	if kerrors.IsNotFound(err) {
		log.Debug("upgrade " + name + " not found")
		return false, err
	}
	if err != nil {
		log.Error("error getting upgrade " + err.Error())
		return false, err
	}

	return true, err
}

// DeleteAllpgupgrade deletes alll pgupgrade
func DeleteAllpgupgrade(client *rest.RESTClient, namespace string) error {

	err := client.Delete().
		Resource(crv1.PgupgradeResourcePlural).
		Namespace(namespace).
		Do().
		Error()
	if err != nil {
		log.Error("error deleting all pgupgrade " + err.Error())
		return err
	}

	log.Debug("deleted all pgupgrade ")
	return err
}

// Deletepgupgrade deletes pgupgrade by name
func Deletepgupgrade(client *rest.RESTClient, name, namespace string) error {

	err := client.Delete().
		Resource(crv1.PgupgradeResourcePlural).
		Namespace(namespace).
		Name(name).
		Do().
		Error()
	if err != nil {
		log.Error("error deleting pgupgrade " + err.Error())
		return err
	}

	log.Debug("deleted pgupgrade " + name)
	return err
}

// Createpgupgrade creates a pgupgrade
func Createpgupgrade(client *rest.RESTClient, upgrade *crv1.Pgupgrade, namespace string) error {

	result := crv1.Pgupgrade{}

	err := client.Post().
		Resource(crv1.PgupgradeResourcePlural).
		Namespace(namespace).
		Body(upgrade).
		Do().
		Into(&result)
	if err != nil {
		log.Error("error creating pgupgrade " + err.Error())
	}

	return err
}
