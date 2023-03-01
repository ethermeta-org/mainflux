package orgs

import (
	"context"

	"github.com/MainfluxLabs/mainflux/auth"
	"github.com/go-kit/kit/endpoint"
)

func createOrgEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createOrgReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		org := auth.Org{
			Name:        req.Name,
			Description: req.Description,
			Metadata:    req.Metadata,
		}

		org, err := svc.CreateOrg(ctx, req.token, org)
		if err != nil {
			return nil, err
		}

		return orgRes{created: true, id: org.ID}, nil
	}
}

func viewOrgEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(orgReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		org, err := svc.ViewOrg(ctx, req.token, req.id)
		if err != nil {
			return nil, err
		}

		res := viewOrgRes{
			ID:          org.ID,
			Name:        org.Name,
			Description: org.Description,
			Metadata:    org.Metadata,
			OwnerID:     org.OwnerID,
			CreatedAt:   org.CreatedAt,
			UpdatedAt:   org.UpdatedAt,
		}

		return res, nil
	}
}

func updateOrgEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateOrgReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		org := auth.Org{
			ID:          req.id,
			Name:        req.Name,
			Description: req.Description,
			Metadata:    req.Metadata,
		}

		_, err := svc.UpdateOrg(ctx, req.token, org)
		if err != nil {
			return nil, err
		}

		return orgRes{created: false}, nil
	}
}

func deleteOrgEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(orgReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.RemoveOrg(ctx, req.token, req.id); err != nil {
			return nil, err
		}

		return deleteRes{}, nil
	}
}

func listOrgsEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listOrgsReq)
		if err := req.validate(); err != nil {
			return nil, err
		}
		pm := auth.PageMetadata{
			Metadata: req.metadata,
		}
		page, err := svc.ListOrgs(ctx, req.token, pm)
		if err != nil {
			return nil, err
		}

		return buildOrgsResponse(page), nil
	}
}

func listMemberships(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listOrgMembershipsReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		pm := auth.PageMetadata{
			Offset:   req.offset,
			Limit:    req.limit,
			Metadata: req.metadata,
		}

		page, err := svc.ListOrgMemberships(ctx, req.token, req.id, pm)
		if err != nil {
			return nil, err
		}

		return buildOrgsResponse(page), nil
	}
}

func assignMembersEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assignMembersReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.AssignMembers(ctx, req.token, req.orgID, req.MemberIDs...); err != nil {
			return nil, err
		}

		return assignRes{}, nil
	}
}

func unassignMembersEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(unassignOrgReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.UnassignMembers(ctx, req.token, req.orgID, req.MemberIDs...); err != nil {
			return nil, err
		}

		return unassignRes{}, nil
	}
}

func listOrgMembersEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listOrgMembersReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		pm := auth.PageMetadata{
			Offset:   req.offset,
			Limit:    req.limit,
			Metadata: req.metadata,
		}
		page, err := svc.ListOrgMembers(ctx, req.token, req.id, pm)
		if err != nil {
			return nil, err
		}

		return buildMembersResponse(page), nil
	}
}

func assignOrgGroupsEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(assignMembersReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.AssignGroups(ctx, req.token, req.orgID, req.MemberIDs...); err != nil {
			return nil, err
		}

		return assignRes{}, nil
	}
}

func unassignOrgGroupsEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(unassignOrgReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.UnassignGroups(ctx, req.token, req.orgID, req.MemberIDs...); err != nil {
			return nil, err
		}

		return unassignRes{}, nil
	}
}

func listOrgGroupsEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listOrgGroupsReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		pm := auth.PageMetadata{
			Offset:   req.offset,
			Limit:    req.limit,
			Metadata: req.metadata,
		}
		page, err := svc.ListOrgGroups(ctx, req.token, req.id, pm)
		if err != nil {
			return nil, err
		}

		return buildGroupsResponse(page), nil
	}
}

func toViewOrgRes(org auth.Org) viewOrgRes {
	view := viewOrgRes{
		ID:          org.ID,
		OwnerID:     org.OwnerID,
		Name:        org.Name,
		Description: org.Description,
		Metadata:    org.Metadata,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	}

	return view
}

func buildOrgsResponse(gp auth.OrgsPage) orgsPageRes {
	res := orgsPageRes{
		pageRes: pageRes{
			Total: gp.Total,
		},
		Orgs: []viewOrgRes{},
	}

	for _, org := range gp.Orgs {
		view := viewOrgRes{
			ID:          org.ID,
			OwnerID:     org.OwnerID,
			Name:        org.Name,
			Description: org.Description,
			Metadata:    org.Metadata,
			CreatedAt:   org.CreatedAt,
			UpdatedAt:   org.UpdatedAt,
		}
		res.Orgs = append(res.Orgs, view)
	}

	return res
}

func buildMembersResponse(mp auth.OrgMembersPage) memberPageRes {
	res := memberPageRes{
		pageRes: pageRes{
			Total:  mp.Total,
			Offset: mp.Offset,
			Limit:  mp.Limit,
			Name:   mp.Name,
		},
		MemberIDs: []string{},
	}

	for _, m := range mp.MemberIDs {
		res.MemberIDs = append(res.MemberIDs, m)
	}

	return res
}

func buildGroupsResponse(mp auth.OrgGroupsPage) groupsPageRes {
	res := groupsPageRes{
		pageRes: pageRes{
			Total:  mp.Total,
			Offset: mp.Offset,
			Limit:  mp.Limit,
			Name:   mp.Name,
		},
		GroupIDs: []string{},
	}

	for _, m := range mp.GroupIDs {
		res.GroupIDs = append(res.GroupIDs, m)
	}

	return res
}