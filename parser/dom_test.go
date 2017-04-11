package parser

import ()

func ExampleOneLine() {
	testdata := "<html><head><title>YO!</title></head><body><h1>lol</h1>dis<elem/></body></html>"
	p := Load(testdata)
	d := p.ParseDOM()
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
