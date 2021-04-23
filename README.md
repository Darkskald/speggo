# speggo

This is a simple server for spectroscopic data (or comparable output of analytical instruments). speggo relies on the gin-gonic web framework.

## Aim

The goal of this project is more of educational purpose and a proof of concept: data that was obtained during a PhD project by different analytical techniques may be served by a REST API rather that can be easily deployed on a web server or in the cloud. Thus, cotainerization of the application will become an issue soon. In general, it is not benefitial that in many cases the everyday data handling of PhD is rather a struggle with distributed data of inconsistent naming over a variety of hard disks, network drives and even measurement computers attached to the corresponding analytica instruments. he deomnstration of the easy of cleaning this mess up employing a golang server is the central idea behind this.

## Future goals

- Serve data of different kinds (at least VSFG spectra, LT compression isotherms, UV/Vis spectra, Raman spetra and infrared spetra)
- query the spectra based on different criteria using query strings
- authorization/ access ccontrol
- clean error handling and intelligent usage of suitable middleware
- add a clear documentation and sketches of the underlying (primtive, tbh) architecture of this API
