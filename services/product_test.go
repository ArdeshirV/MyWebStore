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

func TestProductCRUD(t *testing.T) {
    t.Run("Create Product", func(t *testing.T) {
        for i, product := range testProducts {
            err := productService.Create(product)
            if err != nil {
                AssertNotError(t, err, i)
            }
        }
    })

    t.Run("Create Product Invalid", func(t *testing.T) {
        for i, product := range []*model.Product{
            {},
            {
                Price: 0,
                Weight: 1.5,
                Title: "Samsung Galaxy S24",
            },
            {

                Price: 1000,
                Weight: 0,
                Title: "Samsung Galaxy S24",
            },
            {
                Price: 100,
                Weight: 1.5,
            },
            {
                Price: 1000,
                Weight: 1,
                Title: "Samsung Galaxy S25",
                CategoryID: 1,
                PDF: "file1.pdf",
            },
            {
                Price: 1000,
                Weight: 2,
                Title: "Samsung Galaxy S24",
                CategoryID: 1,
                Description: "Buy me! Buy me now!",
            },
            {
               Price: 1000,
               Weight: 3,
               Title: "Samsung Galaxy S24",
               PDF: "file1.pdf",
               Description: "The best smartphone of 2024",
            },
        } {
            err := productService.Create(product)
            AssertError(t, err, 1)
        }
    })

    t.Run("Update Product", func(t *testing.T) {
        for i, product := range testProducts {
            product.Title += "_UPDATED"
            err := productService.Update(product)
            AssertNotError(t, err, i)
        }
    })

    t.Run("Update Product Invalid", func(t *testing.T) {
        for i, product := range []*models.Product {
            {},
            {ID: 1, Price: 0, Weight: 1.5, Title: "Samsung Galaxy S24"},
            {ID: 2, Price: 1000, Weight: 0, Title: "Samsung Galaxy S24"},
            {ID: 3, Price: 1000, Weight: 1.5, Title: ""},
        } {
            err := productService.Update(product)
            AssertError(t, err, i)
        }
    })

    t.Run("Find Product", func(t *testing.T) {
        product, err := productService.Find(3)
        AssertNotError(t, err, 0)
        if product.ID != 3 {
            t.Fatal("mismatch IDs")
        }
    })

    t.Run("Find Product Invalid", func(t *testing.T) {
        for i, id := range []int{-1, 0, 300} {
            _, err := productService.Find(id)
            AsseertError(t, err, i)
        }
    })

    
}


