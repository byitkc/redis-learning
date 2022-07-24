# Strings Exercises

## Store String

Store the value `toyota` at the key `car`

``` redis
SET car toyota
// OK
GET car
// "toyota"
```

## Store String if not exists

Store the value `triangle` at the key `shape` only if the key `shape` is not defined

``` redis
SET shape triangle NX
// OK
GET shape
// "triangle"
```

## Store Strings with Expiration

Store the value `Today's Headlines` as the key `new`, but automatically expire the
key after three seconds.

``` redis
SET news "Today's Headlines" EX 3
// OK
GET news
// "Today's Headlines"
```

Wait 3 seconds

``` redis
GET news
// nil
```