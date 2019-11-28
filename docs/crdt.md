CRDT
Good resources
https://news.ycombinator.com/item?id=21464189

For anyone more deeply interested in this topic I recommend to read this blog post from Archagon. It describes the different alternatives (OT, CmRDT, CvRDT, diff sync) for writing a collaborative editor. And unlike academic papers it is written in a format how a programmer does research and thinks about a problem in real life, so it's very natural to follow, even if it's long and complex.
http://archagon.net/blog/2018/03/24/data-laced-with-history/

(I am not affiliated in any way, just enjoyed it very much)


	
dboreham 21 days ago [-]

Alexei's work on this subject is very good: the best practical introduction to the subject.
However, I feel that it is worthwhile (necessary?) to completely understand the academic basis for CRDT. A good place to start is Shapiro et al's second paper : https://hal.inria.fr/inria-00555588/document .

In order (sic) to understand the paper you need to have a grasp of Order Theory. This is not terribly hard to get your head around. This is a good place to start: http://jtfmumm.com/blog/2015/11/17/crdt-primer-1-defanging-o... also https://www.wikiwand.com/en/Order_theory and basically stop when you understand this : https://www.wikiwand.com/en/Lattice_(order)

The reason why the formal basis is important is: that's the whole point of CRDT -- previously (been there, got the t-shirt..) folks just made up replication mechanisms they thought would work. Then they build them and embarked on a process of fixing the bugs. Sometimes that took decades. CRDT is nothing new in terms of : software to perform eventually consistent replication. The new thing is that there's a way to formally prove that your bright idea for replication will in fact work (as in : it will converge and it will have the consistency properties you expect). So if you're not seeing that aspect, then you're really not with the program.

btw originally the C stood for Commutative or Convergent, not Conflict-Free. There are plenty of CRDTs that cope with conflicts consistently, rather than being conflict free (e.g. LWW Register).

Projects:
 - Braid HTTP (https://braid.news/)
  - Automerge (https://github.com/automerge/automerge)
  - Gun (https://gun.eco/)
  - Yjs (http://y-js.org/)
  - Noms (https://github.com/attic-labs/noms)
  - DAT (https://dat.foundation/)
  
  
 https://github.com/ipfs/research-CRDT
 
 Order theory
 http://jtfmumm.com/blog/2015/11/17/crdt-primer-1-defanging-order-theory/
 https://www.wikiwand.com/en/Order_theory
 
 History of RIAK using CRDT http://christophermeiklejohn.com/erlang/lasp/2019/03/08/monotonicity.html
 
 
 http://christophermeiklejohn.com/erlang/partisan/2019/04/20/fault-injection-reliable-broadcast.html
 https://managementfromscratch.wordpress.com/2016/04/01/introduction-to-gossip/
