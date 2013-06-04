// groovyc Hello.groovy
// gjars=$(ls /usr/share/groovy/lib/*.jar|fmt -1000|sed 's/ /:/g')
// java -cp .:$gjars Hello

public class Hello {
	public static void main(String[] args) {
		println("Hello in compiled Groovy!")
	}
}
