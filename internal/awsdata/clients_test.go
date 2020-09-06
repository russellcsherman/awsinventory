package awsdata_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var testError = errors.New("test aws error")

type TestClients struct {
	EC2                  ec2iface.EC2API
	ElasticsearchService elasticsearchserviceiface.ElasticsearchServiceAPI
	ELB                  elbiface.ELBAPI
	ELBV2                elbv2iface.ELBV2API
	IAM                  iamiface.IAMAPI
	RDS                  rdsiface.RDSAPI
	Route53              route53iface.Route53API
	S3                   s3iface.S3API
}

func (c TestClients) GetEC2Client(region string) ec2iface.EC2API {
	return c.EC2
}

func (c TestClients) GetIAMClient(region string) iamiface.IAMAPI {
	return c.IAM
}

func (c TestClients) GetElasticsearchServiceClient(region string) elasticsearchserviceiface.ElasticsearchServiceAPI {
	return c.ElasticsearchService
}

func (c TestClients) GetELBClient(region string) elbiface.ELBAPI {
	return c.ELB
}

func (c TestClients) GetELBV2Client(region string) elbv2iface.ELBV2API {
	return c.ELBV2
}

func (c TestClients) GetRDSClient(region string) rdsiface.RDSAPI {
	return c.RDS
}

func (c TestClients) GetRoute53Client(region string) route53iface.Route53API {
	return c.Route53
}

func (c TestClients) GetS3Client(region string) s3iface.S3API {
	return c.S3
}
