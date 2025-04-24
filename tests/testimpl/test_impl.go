package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	awsClient := GetAWSElasticacheClient(t)

	t.Run("TestIsDeployed", func(t *testing.T) {
		parameterGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "elasticache_parameter_group_name")
		out, err := awsClient.DescribeCacheParameterGroups(context.TODO(), &elasticache.DescribeCacheParameterGroupsInput{
			CacheParameterGroupName: aws.String(parameterGroupName),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCacheClusters: %v", err)
		}

		assert.Equal(t, len(out.CacheParameterGroups), 1, "Parameter group name does not exists!")
	})

	t.Run("TestFamily", func(t *testing.T) {
		parameterGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "elasticache_parameter_group_name")
		parameterGroupFamily := terraform.Output(t, ctx.TerratestTerraformOptions(), "elasticache_parameter_group_family")
		out, err := awsClient.DescribeCacheParameterGroups(context.TODO(), &elasticache.DescribeCacheParameterGroupsInput{
			CacheParameterGroupName: aws.String(parameterGroupName),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCacheClusters: %v", err)
		}

		assert.Equal(t, *out.CacheParameterGroups[0].CacheParameterGroupFamily, parameterGroupFamily, "Expected parameter group family did no match current parameter group family!")
	})

	t.Run("TestDeployedParameters", func(t *testing.T) {
		parameterGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "elasticache_parameter_group_name")
		parameterGroupParameters := terraform.OutputListOfObjects(t, ctx.TerratestTerraformOptions(), "elasticache_parameter_group_parameters")

		out, err := awsClient.DescribeCacheParameters(context.TODO(), &elasticache.DescribeCacheParametersInput{
			CacheParameterGroupName: aws.String(parameterGroupName),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCacheClusters: %v", err)
		}

		found := 0

		for _, parameter := range parameterGroupParameters {
			for _, parameter2 := range out.Parameters {
				if parameter["name"].(string) == *parameter2.ParameterName && parameter["value"].(string) == *parameter2.ParameterValue {
					found++
					break
				}
			}
		}

		assert.Equal(t, found, len(parameterGroupParameters), "Expected parameters did not match current parameters!")
	})
}

func GetAWSElasticacheClient(t *testing.T) *elasticache.Client {
	awsElasticacheClient := elasticache.NewFromConfig(GetAWSConfig(t))
	return awsElasticacheClient
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
