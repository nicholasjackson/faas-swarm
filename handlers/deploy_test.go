package handlers

import (
	"testing"

	"github.com/openfaas/faas/gateway/requests"
)

func TestBuildSwarmResourcesAddsMemoryLimits(t *testing.T) {
	req := requests.CreateFunctionRequest{
		Limits: &requests.FunctionResources{
			Memory: "128",
		},
	}

	res := buildResources(&req)

	if res.Limits.MemoryBytes != 128 {
		t.Fatalf("Expected memory limit of 128, got %d", res.Limits.MemoryBytes)
	}
}

func TestBuildSwarmResourcesAddsMemoryReservations(t *testing.T) {
	req := requests.CreateFunctionRequest{
		Requests: &requests.FunctionResources{
			Memory: "128",
		},
		Limits: &requests.FunctionResources{},
	}

	res := buildResources(&req)

	if res.Reservations.MemoryBytes != 128 {
		t.Fatalf("Expected memory limit of 128, got %d", res.Reservations.MemoryBytes)
	}
}

func TestBuildSwarmResourcesWithInvalidMemorySetsReservationsToMinus1(t *testing.T) {
	req := requests.CreateFunctionRequest{
		Requests: &requests.FunctionResources{
			Memory: "wrong",
		},
		Limits: &requests.FunctionResources{},
	}

	res := buildResources(&req)

	if res.Reservations.MemoryBytes != -1 {
		t.Fatalf("Expected memory limit of -1, got %d", res.Reservations.MemoryBytes)
	}
}

func TestBuildSwarmResourcesAddsCPULimits(t *testing.T) {
	req := requests.CreateFunctionRequest{
		Requests: &requests.FunctionResources{},
		Limits: &requests.FunctionResources{
			CPU: "100",
		},
	}

	res := buildResources(&req)

	if res.Limits.NanoCPUs != 100 {
		t.Fatalf("Expected CPU limit of 100, got %d", res.Limits.NanoCPUs)
	}
}

func TestBuildSwarmResourcesAddsCPUReservations(t *testing.T) {
	req := requests.CreateFunctionRequest{
		Requests: &requests.FunctionResources{
			CPU: "100",
		},
		Limits: &requests.FunctionResources{},
	}

	res := buildResources(&req)

	if res.Reservations.NanoCPUs != 100 {
		t.Fatalf("Expected CPU limit of 100, got %d", res.Reservations.NanoCPUs)
	}
}

func TestBuildSwarmResourcesWithInvalidCPUSetsReservationsTo0(t *testing.T) {
	req := requests.CreateFunctionRequest{
		Requests: &requests.FunctionResources{
			Memory: "wrong",
		},
		Limits: &requests.FunctionResources{},
	}

	res := buildResources(&req)

	if res.Reservations.NanoCPUs != -0 {
		t.Fatalf("Expected cpu reservation of 0, got %d", res.Reservations.NanoCPUs)
	}
}
