package tools
 
import (
	_ "github.com/99designs/gqlgen"
)

// mutation create{
//   createShopItem(input: {
//     id: "127",
//     title: "hanakuso",
//     image: "he",
//     description: "hahahahaha"
//   }){
//     title,
//     good,
//     description,
//     seller{
//       id,
//       name
//     }
//   }
// }