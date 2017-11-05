package snmpv2

import (
	"github.com/qmsk/snmpbot/mibs"
	"github.com/qmsk/snmpbot/snmp"
)

var (
	MIB       = mibs.RegisterMIB("SNMPv2-MIB", snmp.OID{1, 3, 6, 1, 6, 3, 1})
	SystemMIB = mibs.RegisterMIB("system", snmp.OID{1, 3, 6, 1, 2, 1, 1})

	SysDescr = SystemMIB.RegisterObject(SystemMIB.Register("sysDescr", 1), mibs.Object{
		Syntax: mibs.DisplayStringSyntax,
	})
	SysObjectID = SystemMIB.RegisterObject(SystemMIB.Register("sysObjectID", 2), mibs.Object{
		Syntax: mibs.ObjectIdentifierSyntax,
	})
	SysUpTime = SystemMIB.RegisterObject(SystemMIB.Register("sysUpTime", 3), mibs.Object{
		Syntax: mibs.TimeTicksSyntax,
	})
	SysContact = SystemMIB.RegisterObject(SystemMIB.Register("sysContact", 4), mibs.Object{
		Syntax: mibs.DisplayStringSyntax,
	})
	SysName = SystemMIB.RegisterObject(SystemMIB.Register("sysName", 5), mibs.Object{
		Syntax: mibs.DisplayStringSyntax,
	})
	SysLocation = SystemMIB.RegisterObject(SystemMIB.Register("sysLocation", 6), mibs.Object{
		Syntax: mibs.DisplayStringSyntax,
	})
	SysServices = SystemMIB.RegisterObject(SystemMIB.Register("sysServices", 7), mibs.Object{
		Syntax: mibs.IntegerSyntax,
	})

	SysORLastChange = SystemMIB.RegisterObject(SystemMIB.Register("sysORLastChange", 8), mibs.Object{
		Syntax: mibs.TimeTicksSyntax,
	})

/*
	SNMPTrapOID = MIB.RegisterOID("snmpTrapOID", 1, 4, 1)

	SNMPv2_coldStart             = SNMPv2MIB.registerNotificationType("coldStart", SNMPv2MIB.define(1, 5, 1)) // 1.3.6.1.6.3.1.1.5.1
	SNMPv2_warmStart             = SNMPv2MIB.registerNotificationType("warmStart", SNMPv2MIB.define(1, 5, 2))
	SNMPv2_authenticationFailure = SNMPv2MIB.registerNotificationType("authenticationFailure", SNMPv2MIB.define(1, 5, 5))
*/
)