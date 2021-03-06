package codedeploy_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/datacratic/aws-sdk-go/aws"
	"github.com/datacratic/aws-sdk-go/aws/awsutil"
	"github.com/datacratic/aws-sdk-go/service/codedeploy"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCodeDeploy_BatchGetApplications() {
	svc := codedeploy.New(nil)

	params := &codedeploy.BatchGetApplicationsInput{
		ApplicationNames: []*string{
			aws.String("ApplicationName"), // Required
			// More values...
		},
	}
	resp, err := svc.BatchGetApplications(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_BatchGetDeployments() {
	svc := codedeploy.New(nil)

	params := &codedeploy.BatchGetDeploymentsInput{
		DeploymentIDs: []*string{
			aws.String("DeploymentId"), // Required
			// More values...
		},
	}
	resp, err := svc.BatchGetDeployments(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_CreateApplication() {
	svc := codedeploy.New(nil)

	params := &codedeploy.CreateApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
	}
	resp, err := svc.CreateApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_CreateDeployment() {
	svc := codedeploy.New(nil)

	params := &codedeploy.CreateDeploymentInput{
		ApplicationName:               aws.String("ApplicationName"), // Required
		DeploymentConfigName:          aws.String("DeploymentConfigName"),
		DeploymentGroupName:           aws.String("DeploymentGroupName"),
		Description:                   aws.String("Description"),
		IgnoreApplicationStopFailures: aws.Boolean(true),
		Revision: &codedeploy.RevisionLocation{
			GitHubLocation: &codedeploy.GitHubLocation{
				CommitID:   aws.String("CommitId"),
				Repository: aws.String("Repository"),
			},
			RevisionType: aws.String("RevisionLocationType"),
			S3Location: &codedeploy.S3Location{
				Bucket:     aws.String("S3Bucket"),
				BundleType: aws.String("BundleType"),
				ETag:       aws.String("ETag"),
				Key:        aws.String("S3Key"),
				Version:    aws.String("VersionId"),
			},
		},
	}
	resp, err := svc.CreateDeployment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_CreateDeploymentConfig() {
	svc := codedeploy.New(nil)

	params := &codedeploy.CreateDeploymentConfigInput{
		DeploymentConfigName: aws.String("DeploymentConfigName"), // Required
		MinimumHealthyHosts: &codedeploy.MinimumHealthyHosts{
			Type:  aws.String("MinimumHealthyHostsType"),
			Value: aws.Long(1),
		},
	}
	resp, err := svc.CreateDeploymentConfig(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_CreateDeploymentGroup() {
	svc := codedeploy.New(nil)

	params := &codedeploy.CreateDeploymentGroupInput{
		ApplicationName:     aws.String("ApplicationName"),     // Required
		DeploymentGroupName: aws.String("DeploymentGroupName"), // Required
		AutoScalingGroups: []*string{
			aws.String("AutoScalingGroupName"), // Required
			// More values...
		},
		DeploymentConfigName: aws.String("DeploymentConfigName"),
		EC2TagFilters: []*codedeploy.EC2TagFilter{
			&codedeploy.EC2TagFilter{ // Required
				Key:   aws.String("Key"),
				Type:  aws.String("EC2TagFilterType"),
				Value: aws.String("Value"),
			},
			// More values...
		},
		ServiceRoleARN: aws.String("Role"),
	}
	resp, err := svc.CreateDeploymentGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_DeleteApplication() {
	svc := codedeploy.New(nil)

	params := &codedeploy.DeleteApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
	}
	resp, err := svc.DeleteApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_DeleteDeploymentConfig() {
	svc := codedeploy.New(nil)

	params := &codedeploy.DeleteDeploymentConfigInput{
		DeploymentConfigName: aws.String("DeploymentConfigName"), // Required
	}
	resp, err := svc.DeleteDeploymentConfig(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_DeleteDeploymentGroup() {
	svc := codedeploy.New(nil)

	params := &codedeploy.DeleteDeploymentGroupInput{
		ApplicationName:     aws.String("ApplicationName"),     // Required
		DeploymentGroupName: aws.String("DeploymentGroupName"), // Required
	}
	resp, err := svc.DeleteDeploymentGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_GetApplication() {
	svc := codedeploy.New(nil)

	params := &codedeploy.GetApplicationInput{
		ApplicationName: aws.String("ApplicationName"), // Required
	}
	resp, err := svc.GetApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_GetApplicationRevision() {
	svc := codedeploy.New(nil)

	params := &codedeploy.GetApplicationRevisionInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Revision: &codedeploy.RevisionLocation{ // Required
			GitHubLocation: &codedeploy.GitHubLocation{
				CommitID:   aws.String("CommitId"),
				Repository: aws.String("Repository"),
			},
			RevisionType: aws.String("RevisionLocationType"),
			S3Location: &codedeploy.S3Location{
				Bucket:     aws.String("S3Bucket"),
				BundleType: aws.String("BundleType"),
				ETag:       aws.String("ETag"),
				Key:        aws.String("S3Key"),
				Version:    aws.String("VersionId"),
			},
		},
	}
	resp, err := svc.GetApplicationRevision(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_GetDeployment() {
	svc := codedeploy.New(nil)

	params := &codedeploy.GetDeploymentInput{
		DeploymentID: aws.String("DeploymentId"), // Required
	}
	resp, err := svc.GetDeployment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_GetDeploymentConfig() {
	svc := codedeploy.New(nil)

	params := &codedeploy.GetDeploymentConfigInput{
		DeploymentConfigName: aws.String("DeploymentConfigName"), // Required
	}
	resp, err := svc.GetDeploymentConfig(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_GetDeploymentGroup() {
	svc := codedeploy.New(nil)

	params := &codedeploy.GetDeploymentGroupInput{
		ApplicationName:     aws.String("ApplicationName"),     // Required
		DeploymentGroupName: aws.String("DeploymentGroupName"), // Required
	}
	resp, err := svc.GetDeploymentGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_GetDeploymentInstance() {
	svc := codedeploy.New(nil)

	params := &codedeploy.GetDeploymentInstanceInput{
		DeploymentID: aws.String("DeploymentId"), // Required
		InstanceID:   aws.String("InstanceId"),   // Required
	}
	resp, err := svc.GetDeploymentInstance(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_ListApplicationRevisions() {
	svc := codedeploy.New(nil)

	params := &codedeploy.ListApplicationRevisionsInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Deployed:        aws.String("ListStateFilterAction"),
		NextToken:       aws.String("NextToken"),
		S3Bucket:        aws.String("S3Bucket"),
		S3KeyPrefix:     aws.String("S3Key"),
		SortBy:          aws.String("ApplicationRevisionSortBy"),
		SortOrder:       aws.String("SortOrder"),
	}
	resp, err := svc.ListApplicationRevisions(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_ListApplications() {
	svc := codedeploy.New(nil)

	params := &codedeploy.ListApplicationsInput{
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListApplications(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_ListDeploymentConfigs() {
	svc := codedeploy.New(nil)

	params := &codedeploy.ListDeploymentConfigsInput{
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListDeploymentConfigs(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_ListDeploymentGroups() {
	svc := codedeploy.New(nil)

	params := &codedeploy.ListDeploymentGroupsInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		NextToken:       aws.String("NextToken"),
	}
	resp, err := svc.ListDeploymentGroups(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_ListDeploymentInstances() {
	svc := codedeploy.New(nil)

	params := &codedeploy.ListDeploymentInstancesInput{
		DeploymentID: aws.String("DeploymentId"), // Required
		InstanceStatusFilter: []*string{
			aws.String("InstanceStatus"), // Required
			// More values...
		},
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListDeploymentInstances(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_ListDeployments() {
	svc := codedeploy.New(nil)

	params := &codedeploy.ListDeploymentsInput{
		ApplicationName: aws.String("ApplicationName"),
		CreateTimeRange: &codedeploy.TimeRange{
			End:   aws.Time(time.Now()),
			Start: aws.Time(time.Now()),
		},
		DeploymentGroupName: aws.String("DeploymentGroupName"),
		IncludeOnlyStatuses: []*string{
			aws.String("DeploymentStatus"), // Required
			// More values...
		},
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListDeployments(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_RegisterApplicationRevision() {
	svc := codedeploy.New(nil)

	params := &codedeploy.RegisterApplicationRevisionInput{
		ApplicationName: aws.String("ApplicationName"), // Required
		Revision: &codedeploy.RevisionLocation{ // Required
			GitHubLocation: &codedeploy.GitHubLocation{
				CommitID:   aws.String("CommitId"),
				Repository: aws.String("Repository"),
			},
			RevisionType: aws.String("RevisionLocationType"),
			S3Location: &codedeploy.S3Location{
				Bucket:     aws.String("S3Bucket"),
				BundleType: aws.String("BundleType"),
				ETag:       aws.String("ETag"),
				Key:        aws.String("S3Key"),
				Version:    aws.String("VersionId"),
			},
		},
		Description: aws.String("Description"),
	}
	resp, err := svc.RegisterApplicationRevision(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_StopDeployment() {
	svc := codedeploy.New(nil)

	params := &codedeploy.StopDeploymentInput{
		DeploymentID: aws.String("DeploymentId"), // Required
	}
	resp, err := svc.StopDeployment(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_UpdateApplication() {
	svc := codedeploy.New(nil)

	params := &codedeploy.UpdateApplicationInput{
		ApplicationName:    aws.String("ApplicationName"),
		NewApplicationName: aws.String("ApplicationName"),
	}
	resp, err := svc.UpdateApplication(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCodeDeploy_UpdateDeploymentGroup() {
	svc := codedeploy.New(nil)

	params := &codedeploy.UpdateDeploymentGroupInput{
		ApplicationName:            aws.String("ApplicationName"),     // Required
		CurrentDeploymentGroupName: aws.String("DeploymentGroupName"), // Required
		AutoScalingGroups: []*string{
			aws.String("AutoScalingGroupName"), // Required
			// More values...
		},
		DeploymentConfigName: aws.String("DeploymentConfigName"),
		EC2TagFilters: []*codedeploy.EC2TagFilter{
			&codedeploy.EC2TagFilter{ // Required
				Key:   aws.String("Key"),
				Type:  aws.String("EC2TagFilterType"),
				Value: aws.String("Value"),
			},
			// More values...
		},
		NewDeploymentGroupName: aws.String("DeploymentGroupName"),
		ServiceRoleARN:         aws.String("Role"),
	}
	resp, err := svc.UpdateDeploymentGroup(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}