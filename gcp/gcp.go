package gcp

import (
	artifactregistry "cloud.google.com/go/artifactregistry/apiv1beta2"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	artifactregistrypb "google.golang.org/genproto/googleapis/devtools/artifactregistry/v1beta2"
)

var ErrNewClient = errors.New("cannot get new client may not have access")

// name "projects/bozo/locations/us-central1"
func Repos(name string) ([]string, error) {
	out := []string{}
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		return out, ErrNewClient
	}
	defer c.Close()

	req := &artifactregistrypb.ListRepositoriesRequest{
		Parent:    name,
		PageSize:  200,
		PageToken: "",
	}

	it := c.ListRepositories(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
			return out, err
		}
		// TODO: Use resp.
		_ = resp
		fmt.Println(resp.Name)
	}
	return out, nil
}

// parent "projects/mchirico/locations/us-central1/repositories/public"
func Files(parent string) ([]string, error) {
	out := []string{}
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		return out, ErrNewClient
	}
	defer c.Close()

	req := &artifactregistrypb.ListFilesRequest{
		Parent:    parent,
		Filter:    "",
		PageSize:  1000,
		PageToken: "",
	}

	it := c.ListFiles(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		out = append(out, resp.Name)

	}

	return out, nil
}

// projects/mchirico/locations/us-central1/repositories/public/packages/septa/tags/v0.0.1
func Delete(name string) error {
	ctx := context.Background()
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
		return ErrNewClient
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteTagRequest{
		Name: name,
	}
	err = c.DeleteTag(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
