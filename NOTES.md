  Hi Balachandar,

You can also look at vboxshell.py, which contains rich set of examples 
on usage of VirtualBox API (and works via both SOAP and COM).
To power off VM code similar to following is used:

from vboxapi import VirtualBoxManager

style = "WEBSERVICE" # or None for local COM/XPCOM
g_virtualBoxManager = VirtualBoxManager(style, None)
mgr=g_virtualBoxManager.mgr
vb=g_virtualBoxManager.vbox
session = mgr.getSessionObject(vb)
uuid = mach.id
try:
progress = vb.openExistingSession(session, uuid)
except Exception,e:
print "Session to '%s' not open: %s" %(mach.name,e)
if g_verbose:
traceback.print_exc()
return
console=session.console
console.powerDown()
session.close()

VirtualBox needs separate session open phase to obtain exclusive control 
over VM execution, that's why sample is a bit more complex that one 
would wish.

Thanks,
Nikolay