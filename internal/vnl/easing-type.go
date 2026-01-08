package vnl

import "github.com/fogleman/ease"

type EasingType string

const (
	Linear       EasingType = "Linear"
	InQuad       EasingType = "InQuad"
	OutQuad      EasingType = "OutQuad"
	InOutQuad    EasingType = "InOutQuad"
	InCubic      EasingType = "InCubic"
	OutCubic     EasingType = "OutCubic"
	InOutCubic   EasingType = "InOutCubic"
	InQuart      EasingType = "InQuart"
	OutQuart     EasingType = "OutQuart"
	InOutQuart   EasingType = "InOutQuart"
	InQuint      EasingType = "InQuint"
	OutQuint     EasingType = "OutQuint"
	InOutQuint   EasingType = "InOutQuint"
	InSine       EasingType = "InSine"
	OutSine      EasingType = "OutSine"
	InOutSine    EasingType = "InOutSine"
	InExpo       EasingType = "InExpo"
	OutExpo      EasingType = "OutExpo"
	InOutExpo    EasingType = "InOutExpo"
	InCirc       EasingType = "InCirc"
	OutCirc      EasingType = "OutCirc"
	InOutCirc    EasingType = "InOutCirc"
	InElastic    EasingType = "InElastic"
	OutElastic   EasingType = "OutElastic"
	InOutElastic EasingType = "InOutElastic"
	InBack       EasingType = "InBack"
	OutBack      EasingType = "OutBack"
	InOutBack    EasingType = "InOutBack"
	InBounce     EasingType = "InBounce"
	OutBounce    EasingType = "OutBounce"
	InOutBounce  EasingType = "InOutBounce"
	InSquare     EasingType = "InSquare"
	OutSquare    EasingType = "OutSquare"
	InOutSquare  EasingType = "InOutSquare"
)

func GetEasingFunc(easingType EasingType) func(float64) float64 {
	switch easingType {
	case Linear:
		return ease.Linear
	case InQuad:
		return ease.InQuad
	case OutQuad:
		return ease.OutQuad
	case InOutQuad:
		return ease.InOutQuad
	case InCubic:
		return ease.InCubic
	case OutCubic:
		return ease.OutCubic
	case InOutCubic:
		return ease.InOutCubic
	case InQuart:
		return ease.InQuart
	case OutQuart:
		return ease.OutQuart
	case InOutQuart:
		return ease.InOutQuart
	case InQuint:
		return ease.InQuint
	case OutQuint:
		return ease.OutQuint
	case InOutQuint:
		return ease.InOutQuint
	case InSine:
		return ease.InSine
	case OutSine:
		return ease.OutSine
	case InOutSine:
		return ease.InOutSine
	case InExpo:
		return ease.InExpo
	case OutExpo:
		return ease.OutExpo
	case InOutExpo:
		return ease.InOutExpo
	case InCirc:
		return ease.InCirc
	case OutCirc:
		return ease.OutCirc
	case InOutCirc:
		return ease.InOutCirc
	case InElastic:
		return ease.InElastic
	case OutElastic:
		return ease.OutElastic
	case InOutElastic:
		return ease.InOutElastic
	case InBack:
		return ease.InBack
	case OutBack:
		return ease.OutBack
	case InOutBack:
		return ease.InOutBack
	case InBounce:
		return ease.InBounce
	case OutBounce:
		return ease.OutBounce
	case InOutBounce:
		return ease.InOutBounce
	case InSquare:
		return ease.InSquare
	case OutSquare:
		return ease.OutSquare
	case InOutSquare:
		return ease.InOutSquare
	}
	return ease.Linear
}
