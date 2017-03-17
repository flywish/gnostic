// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func templates() map[string]string {
	return map[string]string{ 
        "client.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQpwYWNrYWdlIHt7LlJlbmRlcmVyLlBhY2thZ2V9fQoKaW1wb3J0ICgKICAiYnl0ZXMiCiAgImVycm9ycyIKICAiZW5jb2RpbmcvanNvbiIKICAiZm10IgogICJuZXQvaHR0cCIKICAic3RyaW5ncyIKKQogIAovLyBBUEkgY2xpZW50IHJlcHJlc2VudGF0aW9uLgp0eXBlIENsaWVudCBzdHJ1Y3QgewoJc2VydmljZSBzdHJpbmcKfSAKCi8vIENyZWF0ZSBhbiBBUEkgY2xpZW50LgpmdW5jIE5ld0NsaWVudChzZXJ2aWNlIHN0cmluZykgKkNsaWVudCB7CgljbGllbnQgOj0gJkNsaWVudHt9CgljbGllbnQuc2VydmljZSA9IHNlcnZpY2UKCXJldHVybiBjbGllbnQKfQoKLy8te3tyYW5nZSAuUmVuZGVyZXIuTWV0aG9kc319Cnt7Y29tbWVudEZvclRleHQgLkRlc2NyaXB0aW9ufX0KLy8te3tpZiBlcSAuUmVzdWx0VHlwZU5hbWUgIiJ9fQpmdW5jIChjbGllbnQgKkNsaWVudCkge3suQ2xpZW50TmFtZX19KHt7cGFyYW1ldGVyTGlzdCAufX0pIChlcnIgZXJyb3IpIHsKLy8te3tlbHNlfX0KZnVuYyAoY2xpZW50ICpDbGllbnQpIHt7LkNsaWVudE5hbWV9fSh7e3BhcmFtZXRlckxpc3QgLn19KSAocmVzdWx0ICp7ey5SZXN1bHRUeXBlTmFtZX19LCBlcnIgZXJyb3IpIHsKLy8te3tlbmR9fQoJcGF0aCA6PSBjbGllbnQuc2VydmljZSArICJ7ey5QYXRofX0iCgkvLy17e2lmIGhhc1BhcmFtZXRlcnMgLn19CgkvLy17e3JhbmdlIC5QYXJhbWV0ZXJzVHlwZS5GaWVsZHN9fQkKCS8vLXt7aWYgZXEgLlBvc2l0aW9uICJwYXRoIn19CglwYXRoID0gc3RyaW5ncy5SZXBsYWNlKHBhdGgsICJ7IiArICJ7ey5KU09OTmFtZX19IiArICJ9IiwgZm10LlNwcmludGYoIiV2Iiwge3suSlNPTk5hbWV9fSksIDEpCgkvLy17e2VuZH19CgkvLy17e2VuZH19CgkvLy17e2VuZH19CgkvLy17e2lmIGVxIC5NZXRob2QgIlBPU1QifX0KCWJvZHkgOj0gbmV3KGJ5dGVzLkJ1ZmZlcikKCWpzb24uTmV3RW5jb2Rlcihib2R5KS5FbmNvZGUoe3tib2R5UGFyYW1ldGVyTmFtZSAufX0pCglyZXEsIGVyciA6PSBodHRwLk5ld1JlcXVlc3QoInt7Lk1ldGhvZH19IiwgcGF0aCwgYm9keSkKCS8vLXt7ZWxzZX19CglyZXEsIGVyciA6PSBodHRwLk5ld1JlcXVlc3QoInt7Lk1ldGhvZH19IiwgcGF0aCwgbmlsKQoJLy8te3tlbmR9fQoJaWYgZXJyICE9IG5pbCB7CgkJcmV0dXJuCgl9CglyZXNwLCBlcnIgOj0gaHR0cC5EZWZhdWx0Q2xpZW50LkRvKHJlcSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybgoJfQoJaWYgcmVzcC5TdGF0dXNDb2RlID09IDIwMCB7CgkJZGVmZXIgcmVzcC5Cb2R5LkNsb3NlKCkKCQkvLy17e2lmIG5lIC5SZXN1bHRUeXBlTmFtZSAiIn19CgkJZGVjb2RlciA6PSBqc29uLk5ld0RlY29kZXIocmVzcC5Cb2R5KQoJCXJlc3VsdCA9ICZ7ey5SZXN1bHRUeXBlTmFtZX19e30KCQlkZWNvZGVyLkRlY29kZShyZXN1bHQpCgkJLy8te3tlbmR9fQoJfSBlbHNlIHsKCQllcnIgPSBlcnJvcnMuTmV3KHJlc3AuU3RhdHVzKQoJfQoJcmV0dXJuCn0KCi8vLXt7ZW5kfX0KCi8vIHJlZmVyIHRvIGltcG9ydGVkIHBhY2thZ2VzIHRoYXQgbWF5IG9yIG1heSBub3QgYmUgdXNlZCBpbiBnZW5lcmF0ZWQgY29kZQpmdW5jIGZvcmNlZF9wYWNrYWdlX3JlZmVyZW5jZXMoKSB7CglfID0gbmV3KGJ5dGVzLkJ1ZmZlcikKCV8gPSBmbXQuU3ByaW50ZigiIikKCV8gPSBzdHJpbmdzLlNwbGl0KCIiLCIiKQp9",
        "provider.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQpwYWNrYWdlIHt7LlJlbmRlcmVyLlBhY2thZ2V9fQoKLy8gVG8gY3JlYXRlIGEgc2VydmVyLCBmaXJzdCB3cml0ZSBhIGNsYXNzIHRoYXQgaW1wbGVtZW50cyB0aGlzIGludGVyZmFjZS4KLy8gVGhlbiBwYXNzIGFuIGluc3RhbmNlIG9mIGl0IHRvIEluaXRpYWxpemUoKS4KdHlwZSBQcm92aWRlciBpbnRlcmZhY2UgewovLy17e3JhbmdlIC5SZW5kZXJlci5NZXRob2RzfX0KCi8vIFByb3ZpZGVyCnt7Y29tbWVudEZvclRleHQgLkRlc2NyaXB0aW9ufX0KLy8te3tpZiBoYXNQYXJhbWV0ZXJzIC59fQovLy17e2lmIGhhc1Jlc3BvbnNlcyAufX0KICB7ey5Qcm9jZXNzb3JOYW1lfX0ocGFyYW1ldGVycyAqe3suUGFyYW1ldGVyc1R5cGVOYW1lfX0sIHJlc3BvbnNlcyAqe3suUmVzcG9uc2VzVHlwZU5hbWV9fSkgKGVyciBlcnJvcikKLy8te3tlbHNlfX0KICB7ey5Qcm9jZXNzb3JOYW1lfX0ocGFyYW1ldGVycyAqe3suUGFyYW1ldGVyc1R5cGVOYW1lfX0pIChlcnIgZXJyb3IpCi8vLXt7ZW5kfX0KLy8te3tlbHNlfX0KLy8te3tpZiBoYXNSZXNwb25zZXMgLn19CiAge3suUHJvY2Vzc29yTmFtZX19KHJlc3BvbnNlcyAqe3suUmVzcG9uc2VzVHlwZU5hbWV9fSkgKGVyciBlcnJvcikKLy8te3tlbHNlfX0KICB7ey5Qcm9jZXNzb3JOYW1lfX0oKSAoZXJyIGVycm9yKQovLy17e2VuZH19Ci8vLXt7ZW5kfX0JCi8vLXt7ZW5kfX0KfQo=",
        "server.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQpwYWNrYWdlIHt7LlJlbmRlcmVyLlBhY2thZ2V9fQoKaW1wb3J0ICgKCSJlbmNvZGluZy9qc29uIgoJImVycm9ycyIKCSJuZXQvaHR0cCIKCSJzdHJjb252IgoKCSJnaXRodWIuY29tL2dvcmlsbGEvbXV4IgopCgpmdW5jIGludFZhbHVlKHMgc3RyaW5nKSAodiBpbnQ2NCkgewoJdiwgXyA9IHN0cmNvbnYuUGFyc2VJbnQocywgMTAsIDY0KQoJcmV0dXJuIHYKfQoKLy8gVGhpcyBwYWNrYWdlLWdsb2JhbCB2YXJpYWJsZSBob2xkcyB0aGUgdXNlci13cml0dGVuIFByb3ZpZGVyIGZvciBBUEkgc2VydmljZXMuCi8vIFNlZSB0aGUgUHJvdmlkZXIgaW50ZXJmYWNlIGZvciBkZXRhaWxzLgp2YXIgcHJvdmlkZXIgUHJvdmlkZXIKCi8vIFRoZXNlIGhhbmRsZXJzIHNlcnZlIEFQSSBtZXRob2RzLgovLy17e3JhbmdlIC5SZW5kZXJlci5NZXRob2RzfX0KCi8vIEhhbmRsZXIKe3tjb21tZW50Rm9yVGV4dCAuRGVzY3JpcHRpb259fQpmdW5jIHt7LkhhbmRsZXJOYW1lfX0odyBodHRwLlJlc3BvbnNlV3JpdGVyLCByICpodHRwLlJlcXVlc3QpIHsKCXZhciBlcnIgZXJyb3IKCS8vLXt7aWYgaGFzUGFyYW1ldGVycyAufX0KCS8vIGluc3RhbnRpYXRlIHRoZSBwYXJhbWV0ZXJzIHN0cnVjdHVyZQoJdmFyIHBhcmFtZXRlcnMge3suUGFyYW1ldGVyc1R5cGVOYW1lfX0KCS8vLXt7aWYgZXEgLk1ldGhvZCAiUE9TVCJ9fQoJLy8gZGVzZXJpYWxpemUgcmVxdWVzdCBmcm9tIHBvc3QgZGF0YQoJZGVjb2RlciA6PSBqc29uLk5ld0RlY29kZXIoci5Cb2R5KQoJZXJyID0gZGVjb2Rlci5EZWNvZGUoJnBhcmFtZXRlcnMue3tib2R5UGFyYW1ldGVyRmllbGROYW1lIC59fSkKCWlmIGVyciAhPSBuaWwgewoJCXcuV3JpdGVIZWFkZXIoaHR0cC5TdGF0dXNCYWRSZXF1ZXN0KQoJCXcuV3JpdGUoW11ieXRlKGVyci5FcnJvcigpICsgIlxuIikpCgkJcmV0dXJuCgl9CgkvLy17e2VuZH19CgkvLyBnZXQgcmVxdWVzdCBmaWVsZHMgaW4gcGF0aCBhbmQgcXVlcnkgcGFyYW1ldGVycwoJLy8te3tpZiBoYXNQYXRoUGFyYW1ldGVycyAufX0KCXZhcnMgOj0gbXV4LlZhcnMocikKCS8vLXt7ZW5kfX0KCS8vLXt7aWYgaGFzRm9ybVBhcmFtZXRlcnMgLn19CglyLlBhcnNlRm9ybSgpCgkvLy17e2VuZH19CgkvLy17e3JhbmdlIC5QYXJhbWV0ZXJzVHlwZS5GaWVsZHN9fQkKCS8vLXt7aWYgZXEgLlBvc2l0aW9uICJwYXRoIn19CglpZiB2YWx1ZSwgb2sgOj0gdmFyc1sie3suSlNPTk5hbWV9fSJdOyBvayB7CgkJcGFyYW1ldGVycy57ey5OYW1lfX0gPSBpbnRWYWx1ZSh2YWx1ZSkKCX0KCS8vLXt7ZW5kfX0JCgkvLy17e2lmIGVxIC5Qb3NpdGlvbiAiZm9ybWRhdGEifX0KCWlmIGxlbihyLkZvcm1bInt7LkpTT05OYW1lfX0iXSkgPiAwIHsKCQlwYXJhbWV0ZXJzLnt7Lk5hbWV9fSA9IGludFZhbHVlKHIuRm9ybVsie3suSlNPTk5hbWV9fSJdWzBdKQoJfQoJLy8te3tlbmR9fQoJLy8te3tlbmR9fQoJLy8te3tlbmR9fQoJLy8te3tpZiBoYXNSZXNwb25zZXMgLn19CQoJLy8gaW5zdGFudGlhdGUgdGhlIHJlc3BvbnNlcyBzdHJ1Y3R1cmUKCXZhciByZXNwb25zZXMge3suUmVzcG9uc2VzVHlwZU5hbWV9fQoJLy8te3tlbmR9fQoJLy8gY2FsbCB0aGUgc2VydmljZSBwcm92aWRlcgkKCS8vLXt7aWYgaGFzUGFyYW1ldGVycyAufX0KCS8vLXt7aWYgaGFzUmVzcG9uc2VzIC59fQoJZXJyID0gcHJvdmlkZXIue3suUHJvY2Vzc29yTmFtZX19KCZwYXJhbWV0ZXJzLCAmcmVzcG9uc2VzKQoJLy8te3tlbHNlfX0KCWVyciA9IHByb3ZpZGVyLnt7LlByb2Nlc3Nvck5hbWV9fSgmcGFyYW1ldGVycykKCS8vLXt7ZW5kfX0KCS8vLXt7ZWxzZX19CgkvLy17e2lmIGhhc1Jlc3BvbnNlcyAufX0KCWVyciA9IHByb3ZpZGVyLnt7LlByb2Nlc3Nvck5hbWV9fSgmcmVzcG9uc2VzKQoJLy8te3tlbHNlfX0KCWVyciA9IHByb3ZpZGVyLnt7LlByb2Nlc3Nvck5hbWV9fSgpCgkvLy17e2VuZH19CgkvLy17e2VuZH19CQoJaWYgZXJyID09IG5pbCB7CgkvLy17eyBpZiBoYXNSZXNwb25zZXMgLn19CgkJLy8te3sgaWYgLlJlc3BvbnNlc1R5cGUgfCBoYXNGaWVsZE5hbWVkT0sgfX0JCQoJCWlmIHJlc3BvbnNlcy5PSyAhPSBuaWwgewoJCQkvLyB3cml0ZSB0aGUgbm9ybWFsIHJlc3BvbnNlCgkJCWVuY29kZXIgOj0ganNvbi5OZXdFbmNvZGVyKHcpCgkJCWVuY29kZXIuRW5jb2RlKHJlc3BvbnNlcy5PSykKCQkJcmV0dXJuCgkJfSAKCQkvLy17e2VuZH19CgkJLy8te3sgaWYgLlJlc3BvbnNlc1R5cGUgfCBoYXNGaWVsZE5hbWVkRGVmYXVsdCB9fQkJCgkJaWYgcmVzcG9uc2VzLkRlZmF1bHQgIT0gbmlsIHsKCQkJdy5Xcml0ZUhlYWRlcihpbnQocmVzcG9uc2VzLkRlZmF1bHQuQ29kZSkpCgkJCXcuV3JpdGUoW11ieXRlKHJlc3BvbnNlcy5EZWZhdWx0Lk1lc3NhZ2UgKyAiXG4iKSkKCQkJcmV0dXJuCgkJfQoJCS8vLXt7ZW5kfX0KCS8vLXt7ZW5kfX0KCX0gZWxzZSB7CgkJdy5Xcml0ZUhlYWRlcihodHRwLlN0YXR1c0ludGVybmFsU2VydmVyRXJyb3IpCgkJdy5Xcml0ZShbXWJ5dGUoZXJyLkVycm9yKCkgKyAiXG4iKSkKCQlyZXR1cm4KCX0KfQovLy17e2VuZH19CgovLyBJbml0aWFsaXplIHRoZSBBUEkgc2VydmljZS4KZnVuYyBJbml0aWFsaXplKHAgUHJvdmlkZXIpIHsKCXByb3ZpZGVyID0gcAoJdmFyIHJvdXRlciA9IG11eC5OZXdSb3V0ZXIoKXt7cmFuZ2UgLlJlbmRlcmVyLk1ldGhvZHN9fQoJcm91dGVyLkhhbmRsZUZ1bmMoInt7LlBhdGh9fSIsIHt7LkhhbmRsZXJOYW1lfX0pLk1ldGhvZHMoInt7Lk1ldGhvZH19Iil7e2VuZH19CglodHRwLkhhbmRsZSgiLyIsIHJvdXRlcikKfQoKLy8gUHJvdmlkZSB0aGUgQVBJIHNlcnZpY2Ugb3ZlciBIVFRQLgpmdW5jIFNlcnZlSFRUUChhZGRyZXNzIHN0cmluZykgZXJyb3IgewoJaWYgcHJvdmlkZXIgPT0gbmlsIHsKCQlyZXR1cm4gZXJyb3JzLk5ldygiVXNlIHt7LlJlbmRlcmVyLlBhY2thZ2V9fS5Jbml0aWFsaXplKCkgdG8gc2V0IGEgc2VydmljZSBwcm92aWRlci4iKQoJfQoJcmV0dXJuIGh0dHAuTGlzdGVuQW5kU2VydmUoYWRkcmVzcywgbmlsKQp9Cg==",
        "types.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCi8vIFR5cGVzIHVzZWQgYnkgdGhlIEFQSS4KLy8te3tyYW5nZSAuUmVuZGVyZXIuVHlwZXN9fQoKLy8te3tpZiBlcSAuS2luZCAic3RydWN0In19CnR5cGUge3suTmFtZX19IHN0cnVjdCB7IAovLy17e3JhbmdlIC5GaWVsZHN9fQogIHt7Lk5hbWV9fSB7e2dvVHlwZSAuVHlwZX19e3tpZiBuZSAuSlNPTk5hbWUgIiJ9fSBganNvbjoie3suSlNPTk5hbWV9fSJgCi8vLXt7ZW5kfX0KLy8te3tlbmR9fQp9Ci8vLXt7ZWxzZX19CnR5cGUge3suTmFtZX19IHt7LktpbmR9fQovLy17e2VuZH19CgovLy17e2VuZH19",
    }
}