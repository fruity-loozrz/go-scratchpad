api.BPM(100);
api.RPM(33);

const start = 1.7;
const rotationSpeed = 0.5;

function at(t) {
    return start + t * rotationSpeed;
}

api.Action()
    .Dur(0.25)
    .Easing('heavy')
    .Platter(at(0), at(0.1))
    .FaderPattern('open');
api.Action()
    .Dur(0.25)
    .Easing('heavy')
    .Platter(at(0.1), at(0))
    .FaderPattern('open');
api.Action()
    .Dur(0.25)
    .Easing('heavy')
    .Platter(at(0), at(0.1))
    .FaderPattern('open');
api.Action()
    .Dur(0.25)
    .Easing('heavy')
    .Platter(at(0.1), at(0))
    .FaderPattern('open');

api.Action()
    .Dur(0.5)
    .Easing('smooth')
    .Platter(at(0), at(0.4))
    .FaderPattern('open');

api.Action()
    .Dur(0.25)
    .Easing('heavy')
    .Platter(at(0), at(0.1))
    .FaderPattern('open');
api.Action()
    .Dur(0.25)
    .Easing('heavy')
    .Platter(at(0.1), at(0))
    .FaderPattern('open');

api.Action()
    .Dur(0.5)
    .Easing('bounce')
    .Platter(at(0), at(0.4))
    .FaderPattern('open');
api.Action()
    .Dur(0.5)
    .Easing('bounce')
    .Platter(at(0), at(0.4))
    .FaderPattern('open');

api.Action()
    .Dur(0.5)
    .Easing('bounce')
    .Platter(at(0.3), at(0.2))
    .FaderPattern('flare2');

