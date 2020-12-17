<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- ABOUT THE PROJECT -->
## About The Project

Audit-for-Cluster is designed to collect auditing data from the kernel and store them in a centralized Graph Database. It received system call auditing from the ```auditd``` with netlink sockets by using [go-libaudit](https://github.com/elastic/go-libaudit). 

* It is developed with Go
* The audit data is stored in neo4j, easy to query, no need to join mutiple tables
* User can configure auditd with ```auditctl``` as usual and Audit-for-Cluster can comly with changes seamlessly



<p align=center><img src="https://github.com/HanMufu/Audit-for-Cluster/blob/main/Snipaste_2020-12-17_17-43-13.png?raw=true" width="650"/></p>

<!-- USAGE EXAMPLES -->
## Usage

Git clone this repo, add a conf/config.yaml file into your repo as following. 

```
mode: "release"
version: v0.0.1
name: "AuditdForCluster"

neo4j:
  host: "bolt://IP_ADDRESS:PORT"
  user: "neo4j"
  password: "PASSWORD"

log:
  level: "debug"
  filename: "audit-cluster.log"
  max_size: 200
  max_age: 30
  max_backups: 7
```

Use ```auditctl``` set your auditing rules. Run ```go run main.go``` and auditing data will flow into your neo4j database. 

<p align=center><img src="https://github.com/HanMufu/Audit-for-Cluster/blob/main/Snipaste_2020-12-17_17-53-43.png?raw=true" width="700"/></p>


## TODO

In the future, if we have mutiple consumers we can use a Kafka in the middle to handle message forwarding for us. 

<p align=center><img src="https://github.com/HanMufu/Audit-for-Cluster/blob/main/Snipaste_2020-12-17_17-39-19.png?raw=true" width="650"/></p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Mufu Han - [@HanMufu](https://twitter.com/HanMufu) - hanmufu@gmail.com

Project Link: [https://github.com/HanMufu/Audit-for-Cluster](https://github.com/HanMufu/Audit-for-Cluster)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [go-libaudit](https://github.com/elastic/go-libaudit)
* [SPADE](https://github.com/ashish-gehani/SPADE)
* [Ma, Shiqing, et al. "Kernel-supported cost-effective audit logging for causality tracking." 2018 {USENIX} Annual Technical Conference ({USENIX}{ATC} 18). 2018.](https://www.usenix.org/conference/atc18/presentation/ma-shiqing)
* [Ma, Shiqing, et al. "Accurate, low cost and instrumentation-free security audit logging for windows." Proceedings of the 31st Annual Computer Security Applications Conference. 2015.](https://dl.acm.org/doi/abs/10.1145/2818000.2818039)
* [Ma, Shiqing, Xiangyu Zhang, and Dongyan Xu. "Protracer: Towards Practical Provenance Tracing by Alternating Between Logging and Tainting." NDSS. 2016.](https://www.ndss-symposium.org/wp-content/uploads/2017/09/protracer-towards-practical-provenance-tracing-alternating-logging-tainting.pdf)
* [Ma, Shiqing, et al. "{MPI}: Multiple perspective attack investigation with semantic aware execution partitioning." 26th {USENIX} Security Symposium ({USENIX} Security 17). 2017.](https://www.usenix.org/conference/usenixsecurity17/technical-sessions/presentation/ma)
* [Getting started with Linux Audit, SHARE Association, YouTube](https://youtu.be/-BkUGPf0PeQ)
* [golang-neo4j-realworld-example](https://github.com/neo4j-examples/golang-neo4j-realworld-example)
* [Red Hat Enteriprise Linux, 7.6. UNDERSTANDING AUDIT LOG FILES](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/security_guide/sec-understanding_audit_log_files)
* [Check whether a node exists, if not create](https://stackoverflow.com/questions/24015854/check-whether-a-node-exists-if-not-create)
* [auditd-examples](https://github.com/EricGershman/auditd-examples)
* [unistd_64.h](https://android.googlesource.com/platform/prebuilts/gcc/linux-x86/host/x86_64-linux-glibc2.7-4.6/+/refs/heads/jb-dev/sysroot/usr/include/asm/unistd_64.h)





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/HanMufu/Audit-for-Cluster.svg?style=for-the-badge
[contributors-url]: https://github.com/HanMufu/Audit-for-Cluster/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/HanMufu/Audit-for-Cluster.svg?style=for-the-badge
[forks-url]: https://github.com/HanMufu/Audit-for-Cluster/network/members
[stars-shield]: https://img.shields.io/github/stars/HanMufu/Audit-for-Cluster.svg?style=for-the-badge
[stars-url]: https://github.com/HanMufu/Audit-for-Cluster/stargazers
[license-shield]: https://img.shields.io/github/license/HanMufu/Audit-for-Cluster.svg?style=for-the-badge
[license-url]: https://github.com/HanMufu/Audit-for-Cluster/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/mufuhan/
