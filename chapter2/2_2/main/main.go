func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fm: %v\n", err)
			os.Exit(1)
		}
		f := lengconv.Foot(t)
		m := lengconv.Meter(t)
		fmt.Printf("%v = %v, %v = %v\n",
			f, lengconv.FToM(f), m, lengconv.MToF(m))
	}
}
