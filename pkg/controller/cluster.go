package controller

import (
	"github.com/KubeOperator/KubeOperator/pkg/constant"
	"github.com/KubeOperator/KubeOperator/pkg/dto"
	"github.com/KubeOperator/KubeOperator/pkg/service"
	"github.com/kataras/iris/v12/context"
)

type ClusterController struct {
	Ctx                              context.Context
	ClusterService                   service.ClusterService
	ClusterInitService               service.ClusterInitService
	ClusterStorageProvisionerService service.ClusterStorageProvisionerService
	ClusterToolService               service.ClusterToolService
	ClusterNodeService               service.ClusterNodeService
	CLusterLogService                service.ClusterLogService
}

func NewClusterController() *ClusterController {
	return &ClusterController{
		ClusterService:                   service.NewClusterService(),
		ClusterInitService:               service.NewClusterInitService(),
		ClusterStorageProvisionerService: service.NewClusterStorageProvisionerService(),
		ClusterToolService:               service.NewClusterToolService(),
		ClusterNodeService:               service.NewClusterNodeService(),
		CLusterLogService:                service.NewClusterLogService(),
	}
}

func (c ClusterController) Get() (*dto.ClusterPage, error) {
	page, _ := c.Ctx.Values().GetBool("page")
	if page {
		projectName := c.Ctx.URLParam("projectName")
		num, _ := c.Ctx.Values().GetInt(constant.PageNumQueryKey)
		size, _ := c.Ctx.Values().GetInt(constant.PageSizeQueryKey)
		pageItem, err := c.ClusterService.Page(num, size, projectName)
		if err != nil {
			return nil, err
		}
		return &pageItem, nil
	} else {
		var pageItem dto.ClusterPage
		items, err := c.ClusterService.List()
		if err != nil {
			return nil, err
		}
		pageItem.Items = items
		pageItem.Total = len(items)
		return &pageItem, nil
	}
}

func (c ClusterController) GetBy(name string) (*dto.Cluster, error) {
	cl, err := c.ClusterService.Get(name)
	if err != nil {
		return nil, err
	}
	return &cl, nil
}

func (c ClusterController) GetStatusBy(name string) (*dto.ClusterStatus, error) {
	cs, err := c.ClusterService.GetStatus(name)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

func (c ClusterController) Post() (*dto.Cluster, error) {
	var req dto.ClusterCreate
	err := c.Ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}
	item, err := c.ClusterService.Create(req)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c ClusterController) PostInitBy(name string) error {
	return c.ClusterInitService.Init(name)
}

func (c ClusterController) GetProvisionerBy(name string) ([]dto.ClusterStorageProvisioner, error) {
	csp, err := c.ClusterStorageProvisionerService.ListStorageProvisioner(name)
	if err != nil {
		return nil, err
	}
	return csp, nil
}
func (c ClusterController) PostProvisionerBy(name string) (*dto.ClusterStorageProvisioner, error) {
	var req dto.ClusterStorageProvisionerCreation
	err := c.Ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}
	p, err := c.ClusterStorageProvisionerService.CreateStorageProvisioner(name, req)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c ClusterController) DeleteProvisionerBy(clusterName string, name string) error {
	return c.ClusterStorageProvisionerService.DeleteStorageProvisioner(clusterName, name)
}

func (c ClusterController) PostProvisionerBatchBy(clusterName string) error {
	var batch dto.ClusterStorageProvisionerBatch
	if err := c.Ctx.ReadJSON(&batch); err != nil {
		return err
	}
	return c.ClusterStorageProvisionerService.BatchStorageProvisioner(clusterName, batch)
}

func (c ClusterController) GetToolBy(clusterName string) ([]dto.ClusterTool, error) {
	cts, err := c.ClusterToolService.List(clusterName)
	if err != nil {
		return nil, err
	}
	return cts, nil
}

func (c ClusterController) PostToolBy(clusterName string) (*dto.ClusterTool, error) {
	var req dto.ClusterTool
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return nil, err
	}
	cts, err := c.ClusterToolService.Enable(clusterName, req)
	if err != nil {
		return nil, err
	}
	return &cts, nil
}

func (c ClusterController) Delete(name string) error {
	return c.ClusterService.Delete(name)
}

func (c ClusterController) PostBatch() error {
	var batch dto.ClusterBatch
	if err := c.Ctx.ReadJSON(&batch); err != nil {
		return err
	}
	if err := c.ClusterService.Batch(batch); err != nil {
		return err
	}
	return nil
}

func (c ClusterController) GetNodeBy(clusterName string) ([]dto.Node, error) {
	cns, err := c.ClusterNodeService.List(clusterName)
	if err != nil {
		return nil, err
	}
	return cns, nil
}

func (c ClusterController) PostNodeBatchBy(clusterName string) ([]dto.Node, error) {
	var req dto.NodeBatch
	err := c.Ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}
	cns, err := c.ClusterNodeService.Batch(clusterName, req)
	if err != nil {
		return nil, err
	}
	return cns, nil
}

func (c ClusterController) GetWebkubectlBy(clusterName string) (*dto.WebkubectlToken, error) {
	tk, err := c.ClusterService.GetWebkubectlToken(clusterName)
	if err != nil {
		return nil, err
	}

	return &tk, nil
}

func (c ClusterController) GetSecretBy(clusterName string) (*dto.ClusterSecret, error) {
	sec, err := c.ClusterService.GetSecrets(clusterName)
	if err != nil {
		return nil, err
	}
	return &sec, nil
}

func (c ClusterController) GetLogBy(clusterName string) ([]dto.ClusterLog, error) {
	ls, err := c.CLusterLogService.List(clusterName)
	if err != nil {
		return nil, err
	}
	return ls, nil

}
