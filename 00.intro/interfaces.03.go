import "github.com/keradgames/kg-catapult/utils"

type Deployer interface {
	InitDeployment(chan utils.Message) error
	PrepareApplicationBundle() error
	PerformDeploy() error
	DeployStatusChecker() (bool, error)
}

func (bd *BeansTalkDeployer) InitDeployment(messenger chan utils.Message) error {
	/*
		"",
		"Deploy Repository set to: " + bd.GithubRepository,
		"Deploy Branch set to: " + bd.GithubBranch,
		"",
	*/

	return nil
}
