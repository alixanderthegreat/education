require 'minitest/autorun'
require_relative 'hello_world'

class HelloWorldTest < Minitest::Test
  def test_say_hi
    # skip
    assert_equal "Hello, World!", # Commnent
                  HelloWorld.hello
  end
end
