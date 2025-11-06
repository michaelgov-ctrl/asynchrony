This project is a short summary of the concepts that Loris references in the article: https://kristoff.it/blog/asynchrony-is-not-concurrency/

Though it is a forced example that belittles the utility of the Go runtime.

    - Asynchrony: the possibility for tasks to run out of order and still be correct.

    - Concurrency: the ability of a system to progress multiple tasks at a time, be it via parallelism or task switching.

    - Parallelism: the ability of a system to execute more than one task simultaneously at the physical level.

Concurrency is not parallelism context: https://go.dev/blog/waza-talk