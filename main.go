package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func VaultByKey(key string) int{
  for i, v := range db.List {
    if v.Key != key {
      continue
    }
    return i
  }
  return -1
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func MakeKey() string {
  b := make([]rune, 55)
  for i := range b {
    b[i] = chars[rand.Intn(len(chars))]
  }
  return string(b)
}

func main(){
  log.Println("Starting passctl server v2.2")

  ReadConfig()
  LoadDatabase()
  app := fiber.New()

  app.Get("/api/get/:key", func(c *fiber.Ctx) error{
    key := c.Params("key")

    i := VaultByKey(key)
    if i == -1 {
      return c.JSON(&fiber.Map{
        "error": "Invalid key",
     })
    }

    return c.JSON(&fiber.Map{
      "error": "",
      "vault": db.List[i].Data,
    })
  })


  app.Post("/api/set/:key", func(c *fiber.Ctx) error {
    body := struct {
      Vault string `json:"vault"`
    }{}
    key := c.Params("key")

    i := VaultByKey(key)
    if i == -1 {
      return c.JSON(&fiber.Map{
        "error": "Invalid key",
     })
    }

    if cfg.MaxVaultSize != 0 && len(c.Body())/1048576 >= cfg.MaxVaultSize {
      return c.JSON(&fiber.Map{
        "error": fmt.Sprintf("Reached max vault size (%d MB)", cfg.MaxVaultSize),
      })
    }

    c.BodyParser(&body)
    db.List[i].Data = body.Vault 
    SaveDatabase()

    return c.JSON(&fiber.Map{
      "error": "",
    })
  })

  app.Get("/api/ping/:key", func(c *fiber.Ctx) error {
    key := c.Params("key")
    
    i := VaultByKey(key)
    if i == -1 {
      return c.JSON(&fiber.Map{
        "error": "Invalid key",
     })
    }

    return c.JSON(&fiber.Map{
      "error": "",
    })
  })

  app.Get("/api/password", func(c *fiber.Ctx) error {
    return c.JSON(&fiber.Map{
      "enabled": cfg.Password != "",
    })
  })

  app.Get("/api/gen", func(c *fiber.Ctx) error {
    pwd := c.Query("password")
    if pwd != cfg.Password {
      return c.JSON(&fiber.Map{
        "error": "Bad password",
      })
    }

    if cfg.MaxVaultCount != 0 && len(db.List) >= cfg.MaxVaultCount {
      log.Printf("Rejecting generation, hit max limit (%d/%d)", len(db.List), cfg.MaxVaultCount)
      return c.JSON(&fiber.Map{
        "error": "Server reached max vault limit",
      })
    }

    key := MakeKey()
    db.List = append(db.List, Vault{
      Key: key,
      Data: "",
    })
    SaveDatabase()

    return c.JSON(&fiber.Map{
      "error": "",
      "key": key, 
    })
  })

  app.Static("/", "./public")
  log.Fatal(app.Listen(cfg.Port))
}
