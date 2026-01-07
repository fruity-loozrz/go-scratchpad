package vnl

import "github.com/fogleman/ease"

func GetEasingFunc(easingType EasingType) func(float64) float64 {
	switch easingType {
	case EaseLinear:
		return ease.Linear
	case EaseSmooth:
		return ease.InOutSine
	case EaseSharp:
		return ease.InOutCubic
	case EaseHeavy:
		return ease.InOutQuint
	case EaseMotorStart:
		return ease.OutExpo
	case EaseSpinback:
		return ease.OutCubic
	case EasePowerDown:
		return ease.OutQuad
	case EaseBounce:
		return ease.OutBounce
	case EaseElastic:
		return ease.OutElastic
	}
	return ease.Linear
}
