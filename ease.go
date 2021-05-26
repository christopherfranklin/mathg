package mathg

import "math"

// Useful functions for physics/graphics

func QuadraticEaseOut(f float64) float64 {
	return -f * (f - 2.)
}

func QuadraticEaseIn(f float64) float64 {
	return f * f
}

func QuadraticEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 2. * f * f
	} else {
		a = -2.*f*f + 4.*f - 1.
	}
	return a
}

func CubicEaseOut(f float64) float64 {
	a := f - 1.
	return a*a*a + 1.
}

func CubicEaseIn(f float64) float64 {
	return f * f * f
}

func CubicEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 4. * f * f * f
	} else {
		a = 2.*f - 2.
		a = 0.5*a*a*a + 1.
	}
	return a
}

func QuarticEaseOut(f float64) float64 {
	a := f - 1.
	return a*a*a*(1.-f) + 1.
}

func QuarticEaseIn(f float64) float64 {
	return f * f * f * f
}

func QuarticEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 8. * f * f * f * f
	} else {
		a = f - 1.
		a = -8.*a*a*a*a + 1.
	}
	return a
}

func QuinticEaseOut(f float64) float64 {
	a := f - 1.
	return a*a*a*a*a + 1.
}

func QuinticEaseIn(f float64) float64 {
	return f * f * f * f * f
}

func QuinticEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 8. * f * f * f * f * f
	} else {
		a = 2.*f - 2.
		a = 0.5*a*a*a*a*a + 1.
	}
	return a
}

func SinEaseOut(f float64) float64 {
	return math.Sin(f * pi_2)
}

func SinEaseIn(f float64) float64 {
	return math.Sin((f-1.)*pi_2) + 1.
}

func SinEaseInOut(f float64) float64 {
	return 0.5 * (1. - math.Cos(f*math.Pi))
}

func CircularEaseOut(f float64) float64 {
	return math.Sqrt((2. - f) * f)
}

func CircularEaseIn(f float64) float64 {
	return 1. - math.Sqrt((1. - (f * f)))
}

func CircularEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 0.5 * (1. - math.Sqrt(1.-4.*f*f))
	} else {
		a = 0.5 * (math.Sqrt(-2.*f-3.)*(2.*f-1.) + 1.)
	}
	return a
}

func ExponentialEaseOut(f float64) float64 {
	a := f
	if math.Abs(a) > epsilon {
		a = 1. - math.Pow(2., -10.*f)
	}
	return a
}

func ExponentialEaseIn(f float64) float64 {
	a := f
	if math.Abs(a) > epsilon {
		a = math.Pow(2., 10.*(f-1.))
	}
	return a
}

func ExponentialEaseInOut(f float64) float64 {
	a := f
	if f < 0.5 {
		a = 0.5 * math.Pow(2., (20.*f)-10.)
	} else {
		a = -0.5*math.Pow(2., -20.*f+10.) + 1.
	}
	return a
}

func ElasticEaseOut(f float64) float64 {
	return math.Sin(-13.*pi_2*(f+1.))*math.Pow(2., -10.*f) + 1.
}

func ElasticEaseIn(f float64) float64 {
	return math.Sin(13.*pi_2*f) * math.Pow(2., 10.*(f-1.))
}

func ElasticEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 0.5 * math.Sin(13.*pi_2*(2.*f)) * math.Pow(2., 10.*((2.*f)-1.))
	} else {
		a = 0.5 * (math.Sin(-13.*pi_2*((2.*f-1.)+1.))*math.Pow(2., -10.*(2.*f-1.)) + 2.)
	}
	return a
}

func BackEaseOut(f float64) float64 {
	a := 1. - f
	return 1. - (a*a*a - a*math.Sin(a*math.Pi))
}

func BackEaseIn(f float64) float64 {
	return f*f*f - f*math.Sin(f*math.Pi)
}

func BackEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 2. * f
		a = 0.5 * (a*a*a - a*math.Sin(a*math.Pi))
	} else {
		a = 1. - (2.*f - 1.)
		a = 0.5*(1.-(a*a*a-a*math.Sin(f*math.Pi))) + 0.5
	}
	return a
}

func BounceEaseOut(f float64) float64 {
	a := 0.
	if f < 4./11. {
		a = (121. * f * f) / 16.
	} else if f < 8./11. {
		a = (363. / 40. * f * f) - (99. / 10. * f) + 17./5.
	} else if f < 9./10. {
		a = (4356. / 361. * f * f) - (513. / 25. * f) + 268./25.
	} else {
		a = (54. / 5. * f * f) - (513. / 25. * f) + 268./25.
	}
	return a
}

func BounceEaseIn(f float64) float64 {
	return 1. - BounceEaseOut(1.-f)
}

func BounceEaseInOut(f float64) float64 {
	a := 0.
	if f < 0.5 {
		a = 0.5 * BounceEaseIn(f*2.)
	} else {
		a = 0.5*BounceEaseOut(f*2.-1.) + 0.5
	}
	return a
}
