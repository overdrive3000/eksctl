package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSRoboMakerRobotApplicationVersion AWS CloudFormation Resource (AWS::RoboMaker::RobotApplicationVersion)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html
type AWSRoboMakerRobotApplicationVersion struct {

	// Application AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html#cfn-robomaker-robotapplicationversion-application
	Application *Value `json:"Application,omitempty"`

	// CurrentRevisionId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html#cfn-robomaker-robotapplicationversion-currentrevisionid
	CurrentRevisionId *Value `json:"CurrentRevisionId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoboMakerRobotApplicationVersion) AWSCloudFormationType() string {
	return "AWS::RoboMaker::RobotApplicationVersion"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSRoboMakerRobotApplicationVersion) MarshalJSON() ([]byte, error) {
	type Properties AWSRoboMakerRobotApplicationVersion
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSRoboMakerRobotApplicationVersion) UnmarshalJSON(b []byte) error {
	type Properties AWSRoboMakerRobotApplicationVersion
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSRoboMakerRobotApplicationVersion(*res.Properties)
	}

	return nil
}

// GetAllAWSRoboMakerRobotApplicationVersionResources retrieves all AWSRoboMakerRobotApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSRoboMakerRobotApplicationVersionResources() map[string]AWSRoboMakerRobotApplicationVersion {
	results := map[string]AWSRoboMakerRobotApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSRoboMakerRobotApplicationVersion:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::RoboMaker::RobotApplicationVersion" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSRoboMakerRobotApplicationVersion{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSRoboMakerRobotApplicationVersionWithName retrieves all AWSRoboMakerRobotApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRoboMakerRobotApplicationVersionWithName(name string) (AWSRoboMakerRobotApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSRoboMakerRobotApplicationVersion:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::RoboMaker::RobotApplicationVersion" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSRoboMakerRobotApplicationVersion{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSRoboMakerRobotApplicationVersion{}, errors.New("resource not found")
}
