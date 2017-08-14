package resources

// AWS::EC2::Instance.SsmAssociation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html
type AWSEC2InstanceSsmAssociation struct {

	// AssociationParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-associationparameters
	AssociationParameters []AWSEC2InstanceSsmAssociationAssociationParameter `json:"AssociationParameters"`

	// DocumentName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html#cfn-ec2-instance-ssmassociations-documentname
	DocumentName string `json:"DocumentName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InstanceSsmAssociation) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.SsmAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InstanceSsmAssociation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}