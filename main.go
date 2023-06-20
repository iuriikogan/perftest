package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/iuriikogan/perftest/imports/k8s"
)

type PerfTestChartProps struct {
	cdk8s.ChartProps
}

func NewPerfTestChart(scope constructs.Construct, id string, props *PerfTestChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	label := map[string]*string{"app": jsii.String("nginx")}

	k8s.NewKubeService(chart, jsii.String("service"), &k8s.KubeServiceProps{
		Spec: &k8s.ServiceSpec{
			Type: jsii.String("LoadBalancer"),
			Ports: &[]*k8s.ServicePort{{
				Port:       jsii.Number(80),
				TargetPort: k8s.IntOrString_FromNumber(jsii.Number(8080)),
			}},
			Selector: &label,
		},
	})

	k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
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
						Name:  jsii.String("nginx"),
						Image: jsii.String("nginx:1.7.9"),
						Ports: &[]*k8s.ContainerPort{{ContainerPort: jsii.Number(8080)}},
					}},
				},
			},
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewPerfTestChart(app, "nginx", nil)
	app.Synth()
}
