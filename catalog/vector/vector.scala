// mutable

abstract class AbstractVector[A] {
  var as = new Array[Any](10)
  var size = 0
  var copyCount = 0

  def add(a: A):Unit = {
    if (size == as.length) {
      resize()
    }
    as(size) = a
    size = size + 1
    copyCount = copyCount + 1
  }

  def apply(i: Int): A = {
    if (size <= i) throw new ArrayIndexOutOfBoundsException
    as(i).asInstanceOf[A]
  }

  def first: A = {
    as(0).asInstanceOf[A]
  }

  def last: A = {
    as(size - 1).asInstanceOf[A]
  }

  override def toString: String = {
    "[" + as.filter(_ != null).mkString(",") + "]"
  }

  protected def resize(): Unit = {
    val bs = as
    as = newArray(as)
    size = 0
    for (x <- bs) add(x.asInstanceOf[A])
  }

  protected def newArray(as: Array[Any]): Array[Any]
}

class BetterVector[A] extends AbstractVector[A] {
  override protected def newArray(as: Array[Any]): Array[Any] = {
    new Array[Any](as.length * 2)
  }
}

class BadVector[A] extends AbstractVector[A] {
  override protected def newArray(as: Array[Any]): Array[Any] = {
    new Array[Any](as.length + 10)
  }
}
