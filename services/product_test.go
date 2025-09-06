package services_test

import (
    "testing"
    "github.com/ArdeshirV/MyWebStore/models"
    "github.com/ArdeshirV/MyWebStore/services"
    "github.com/ArdeshirV/MyWebStore/storage"
)

func AssertNotError(t *testing.T, err error, i int) {
    if err != nil {
        return
    }
    t.Fatalf("[test case %d] expected nil got %v", i, err)
}

func AssertError(t *testing.T, err error, i int) {
    if err != nil {
        return
    }
    t.Fatalf("[test case %d] expected error got nil", i)
}

var (
    productStorage = storage.NewProductInMemory()
    productService = services.NewProduct(productStorage)
    testProducts = []*models.Product{
        {
            Price: 1000,
            Weight: 1.5,
            Title: "Samsung Galaxy S24",
            CategoryID: 1,
            PDF: "file1.pdf",
            Description: "Perfect product!",
        },
        {
            Price: 2000,
            Weight: 5.7,
            Title: "Playstation 5 Pro",
            CategoryID: 2,
            PDF: "file2.pdf",
            Description: "Buy it now!",
        },
        {
            Price: 3000,
            Weight: 2.1,
            Title: "Galaxt Buds Pro",
            CategoryID: 3,
            PDF: "file3.pdf",
            Description: "The best price in the market!",
        },
    }
)


