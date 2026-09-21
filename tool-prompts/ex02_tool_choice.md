## ex:
get_weather(location)        -> weather for a place
define_word(word)            -> a dictionary definition
get_statistic(country, metric) -> value of a given metric for a country 
## 1: Reply with only a JSON object naming the tool and its arguments, like {"tool": ..., "args": {...}}.

## 2: what's it like in Paris? && what does 'recursion' mean?

`outpuy:`
{"tool": "get_weather", "args": {"location": "Paris"}}

{"tool": "define_word", "args": {"word": "recursion"}}

## 3: did the model pick the right tool each time,
Yes both correct:
##  and what is your code's job after it replies?
Parse the JSON response.
Look up the tool name in a real dictionary/map of actual functions you wrote (e.g. {"get_weather": get_weather_fn, "define_word": define_word_fn}).
Call that function, unpacking args into it - e.g. tools[response["tool"]](**response["args"]).
Do something with the real result (actual weather data, actual definition) - not what the model said, since the model doesn't actually know today's weather, it just decided which tool should.