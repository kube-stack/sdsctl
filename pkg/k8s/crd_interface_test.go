package k8s

import (
	"context"
	"fmt"
	"github.com/kube-stack/sdsctl/pkg/constant"
	"github.com/tidwall/gjson"
	"testing"
)

func TestGet2(t *testing.T) {
	ksgvr := NewKsGvr(constant.VMDS_Kind)
	vmd, err := ksgvr.Get(context.TODO(), "default", "disktest")
	if err != nil {
		panic(err)
	}
	fmt.Printf("disktest131: %+v\n", vmd)
	fmt.Printf("disktest131 spec: %+v\n", string(vmd.Spec.Raw))

	parse := gjson.ParseBytes(vmd.Spec.Raw)
	nodeName := parse.Get("nodeName")
	msg := parse.Get("status.conditions.state.waiting.message")
	fmt.Printf("disktest131 spec nodename: %+v\n", nodeName)
	fmt.Printf("disktest131 spec msg: %+v\n", msg)
}

func TestGet3(t *testing.T) {
	ksgvr := NewKsGvr(constant.VMPS_Kind)
	vmp, err := ksgvr.Get(context.TODO(), "default", "pooltest")
	if err != nil {
		panic(err)
	}
	fmt.Printf("pooltest: %+v\n", vmp)
	fmt.Printf("pooltest spec: %+v\n", string(vmp.Spec.Raw))

	parse := gjson.ParseBytes(vmp.Spec.Raw)
	nodeName := parse.Get("nodeName")
	msg := parse.Get("status.conditions.state.waiting.message")
	fmt.Printf("pooltest spec nodename: %+v\n", nodeName)
	fmt.Printf("pooltest spec msg: %+v\n", msg)
}

func TestKsGvr_Update(t *testing.T) {
	ksgvr := NewKsGvr(constant.VMPS_Kind)
	data := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]string{
			"key3": "value3",
		},
	}
	err := ksgvr.Update(context.TODO(), "default", "pooltest111", "pool.extra", data)
	fmt.Printf("err: %+v\n", err)
}

func TestKsGvr_Delete(t *testing.T) {
	ksgvr := NewKsGvr(constant.VMPS_Kind)
	err := ksgvr.Delete(context.TODO(), "default", "pooltest111")
	fmt.Printf("err: %+v\n", err)
}

func TestKsGvr_Create(t *testing.T) {
	ksgvr := NewKsGvr(constant.VMDS_Kind)
	res := make(map[string]string)
	res["key1"] = "value1"
	res["key2"] = "value2"
	err := ksgvr.Create(context.TODO(), constant.DefaultNamespace, "disktest222", constant.CRD_Volume_Key, res)
	fmt.Printf("err: %+v\n", err)
}

func TestKsGvr_CreateExternalCrd(t *testing.T) {
	ksgvr := NewExternalGvr(constant.DefaultRookGroup, constant.DefaultRookVersion, constant.CephBlockPoolS_Kinds)
	res := make(map[string]interface{})
	res["failureDomain"] = "host"
	res["replicated"] = map[string]interface{}{
		"size":                   3,
		"requireSafeReplicaSize": true,
	}
	err := ksgvr.CreateExternalCrd(context.TODO(), constant.RookNamespace, "rbdpooltest", "spec", res)
	fmt.Printf("err: %+v\n", err)
}
