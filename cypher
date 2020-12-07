CREATE (
   SYSCALL3132:SYSCALL
   {    
      arch:"c000003e",
      pid:3132,
      syscall:228,
      succes:"yes",
      exit:0,
      a0:7,
      a1:"7ffeaa82f500",
      a2:"55a7c0776cf0",
      a3:0,
      items:0,
      ppid:3133,
      auid:1000,
      uid:1000,
      gid:1001,
      euid:1000,
      suid:1000,
      fsuid:1000,
      egid:1001,
      sgid:1001,
      fsgid:1001,
      tty:"(none)",
      ses:6,
   comm:"sshd",
   exe:"/usr/sbin/sshd",
   key:"(null)"     
   }
);

CREATE (
   SYSCALL3133:SYSCALL
   {    
      arch:"c000003e",
      pid:3133,
      syscall:228,
      succes:"yes",
      exit:0,
      a0:7,
      a1:"7ffeaa82f500",
      a2:"55a7c0776cf0",
      a3:0,
      items:0,
      ppid:1893,
      auid:1000,
      uid:1000,
      gid:1001,
      euid:1000,
      suid:1000,
      fsuid:1000,
      egid:1001,
      sgid:1001,
      fsgid:1001,
      tty:"(none)",
      ses:6,
   comm:"sshd",
   exe:"/usr/sbin/sshd",
   key:"(null)"     
   }
);


CREATE (
   SYSUSER1000:SYSUSER
   {    
     uid:1000    
   }
);
CREATE (
   GROUP1001:GROUP
   {    
     gid:1001    
   }
);

CREATE (SYSCALL1822)-[:pp]->(SYSCALL1816);
CREATE (SYSCALL1822)-[:au]->(SYSUSER1000);
CREATE (SYSCALL1822)-[:g]->(GROUP1001);
CALL db.relationshipTypes()

match (a:SYSCALL),(b:SYSCALL)
where a.pid='1822' and b.uid='1000'
create (a)-[:pp]->(b);

match (a:SYSCALL),(b:SYSUSER)
where a.pid='1822' and b.uid='1000'
create (a)-[:au]->(b);

match (a:SYSCALL),(b:GROUP)
where a.pid='1822' and b.gid='1001'
create (a)-[:g]->(b);

match (s:SYSCALL),(u:SYSUSER),(g:GROUP) return s,u,g;

MATCH (a:SYSCALL)
WHERE a.pid = '13977'
RETURN a limit 1