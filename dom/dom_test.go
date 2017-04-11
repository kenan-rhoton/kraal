package dom

func ExampleNode_PrettyPrint() {
	parent := EmptyElem("parent")
	parent.Append(EmptyElem("child"))
	parent.Child(0).Append(EmptyElem("grandchild"))
	parent.Child(0).Append(Text("This is simply text"))
	parent.Append(EmptyElem("child"))
	parent.PrettyPrint()
	// Output:
	// <parent>
	//   <child>
	//     <grandchild/>
	//     This is simply text
	//   </child>
	//   <child/>
	// </parent>
}
