package main

import (
	"fmt"

	"gihub.com/iuriikogan/perftest/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String("deployment"), &cprops)

	label := map[string]*string{"app": jsii.String("hello-k8s")}

	obj := &k8s.KubeDeploymentProps{
		Spec: &k8s.DeploymentSpec{
			Replicas: jsii.Number(2),
			Selector: &k8s.LabelSelector{
				MatchLabels: &label,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &label,
				},
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{{
						Name:  jsii.String("GoDeployment"),
						Image: jsii.String("GoDeployment:1.0.0"),
						Ports: &[]*k8s.ContainerPort{{ContainerPort: jsii.Number(8080)}},
					}},
				},
			},
		},
	}

	k8s.NewKubeDeployment(chart, jsii.String("goDeployment"), obj)

	fmt.Println(*(*(obj.Spec.Template.Spec.Containers))[0].Image)

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "cdk8sgo", nil)
	app.Synth()
}
