package parser

import ()

func ExampleOneLine() {
	testdata := "<html><head><title>YO!</title></head><body><h1>lol</h1>dis<elem/></body></html>"
	d := ParseDOM(testdata)
	d.PrettyPrint()
	// Output:
	// <document>
	//   <html>
	//     <head>
	//       <title>
	//         YO!
	//       </title>
	//     </head>
	//     <body>
	//       <h1>
	//         lol
	//       </h1>
	//       dis
	//       <elem/>
	//     </body>
	//   </html>
	// </document>
}

func ExampleSpaced() {
	testdata := "   \t  <html>\n\n \t<head><title>     YO!</title></head><body   ><h1>lol</h1>dis<elem /></body></html>                     "
	d := ParseDOM(testdata)
	d.PrettyPrint()
	// Output:
	// <document>
	//   <html>
	//     <head>
	//       <title>
	//         YO!
	//       </title>
	//     </head>
	//     <body>
	//       <h1>
	//         lol
	//       </h1>
	//       dis
	//       <elem/>
	//     </body>
	//   </html>
	// </document>
}
