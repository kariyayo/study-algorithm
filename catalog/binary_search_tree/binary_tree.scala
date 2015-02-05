object TreeModule {

  sealed trait Tree[A] {
    def insert(x: A): Tree[A]
    def contains(x: A): Boolean
    def remove(x: A): Tree[A]
  }

  case class Empty[A <% Ordered[A]]() extends Tree[A] {
    override def insert(x: A) = Node(x, Empty[A](), Empty[A]())
    override def contains(x: A) = false
    override def remove(x: A) = Empty[A]()
  }

  case class Node[A <% Ordered[A]](a: A, left: Tree[A], right: Tree[A]) extends Tree[A] {
    override def insert(x: A): Tree[A] =
      if (x == a)     Node(a, left, right)
      else if (x < a) Node(a, left.insert(x), right)
      else            Node(a, left, right.insert(x))

    override def contains(x: A): Boolean =
      if (x == a)     true
      else if (x < a) left.contains(x)
      else            right.contains(x)

    override def remove(x: A) = (left, right) match {
      case (_: Empty[A], _: Empty[A]) =>
        if (x == a) Empty[A]()
        else        Node(a, Empty[A](), Empty[A]())
      case (l: Node[A], _: Empty[A]) =>
        if (x == a) l
        else        Node(a, l.remove(x), Empty[A]())
      case (_: Empty[A], r: Node[A]) =>
        if (x == a) r
        else        Node(a, Empty[A](), r.remove(x))
      case (l: Node[A], r: Node[A]) => {
        val maxFromRight = maxValue(r)
        if (x == a)     Node(maxFromRight, l, r.remove(maxFromRight))
        else if (x < a) Node(a, l.remove(x), r)
        else            Node(a, l, r.remove(x))
      }
    }

    private def maxValue(node: Node[A]): A = node match {
      case Node(v, _: Empty[A], _: Empty[A]) => v
      case Node(v, l: Node[A], _: Empty[A])  => List(v, maxValue(l)).max
      case Node(v, _: Empty[A], r: Node[A])  => List(v, maxValue(r)).max
      case Node(v, l: Node[A], r: Node[A])   => List(v, maxValue(l), maxValue(r)).max
    }
  }

}

