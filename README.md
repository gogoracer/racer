# racer
# Go go racer!!!  
gotta go fast

![](docs/dendron/notes/assets/2023-01-27-21-54-23.png)

## What
Type safe web frontends built in Go.  Its like frontend rails but faster.

## Are you crazy
Yes.  Crazy of fighting the churn on the frontend.  And guess what, there big push now is SSR (Serve Side Rendering) with islands of reactivity. [Like](https://qwik.builder.io/docs/overview/) [every](https://astro.build/) [meta](https://nuxt.com/) [framework](https://nextjs.org/) [released](https://kit.svelte.dev/).  Guess what's REALLY good and server side web dev... GOPHERS!

Now imagine if you could teach a gopher to ride a cafe racer custom built for his cute little frame, [sanic](https://knowyourmeme.com/memes/sanic-hegehog) move over; we have a new kid on the block.

## Cool speech bro, what are you really proposing

```
Add diagram here of SPA vs SSR vs LiveView
```
We stand on the shoulders of [work](https://hexdocs.pm/phoenix_live_view/Phoenix.LiveView.html) [done](https://github.com/jfyne/live) [before](https://github.com/SamHennessy/hlive) and consolidate into a Realtime smart framework that out of the box has

1. Completely Go based programming model.  You can still use JS/TS, but never **have** to.
1. Type safe DOM.  It's our contention you'll be FASTER than you would be in React/Vue/Svelte/Solid.  Not just at runtime, but in development
1. Faster paints, our overhead is currents **7Kb**.  That's really small, look at the graph below.  First paint faster
1. Smart real world defaults.  
   1. Auto-reloading
   1. Auto reconnect with exponential backoff
   1. Bundler
1. Its a framework without the headache, want to replace a part, no problem.  Its interfaces at the higher level.  The normal fears about code size do apply here, its all server side
   

```
ADD Graphs of size and first paint
```

## How do I get started?  

We are deep in the conversion to a full stack framework so check back or dig into the code and start asking questions on the #gopher slack

## Got docs/notes about the design process
Yes!  We are using [Dendron](https://wiki.dendron.so/) for our note taking.  It's markdown on steroids and even if you don't use VSCode for programming its a powerful for tracking ideas across a team.