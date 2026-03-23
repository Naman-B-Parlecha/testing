# Agentic Skill Design Patterns

A guide to building skills that produce autonomous, strategic agent behavior — not rigid tool-call scripts.

---

## The Core Problem

Most skills are written like recipes:

```markdown
1. Call tool A
2. Pass result to tool B
3. Format output
```

This creates a **brittle pipeline**. The agent can't recover from failures, can't adapt to unexpected inputs, and can't make judgment calls. It's not an agent — it's a script with extra steps.

The fix: **define the problem space, not the solution steps.**

---

## The Five Layers of Agentic Skill Design

### Layer 1: Goal Definition (the "what")

State the outcome, not the process. The agent needs to know what success looks like so it can evaluate its own progress and course-correct.

```markdown
# BAD — procedural
## Steps
1. Search for candidates matching the job description
2. Enrich top 10 results
3. Score each candidate
4. Return ranked list

# GOOD — goal-oriented
## Objective
Produce a ranked shortlist of candidates who match the hiring criteria,
with enough profile data to make an outreach decision.

## Success criteria
- Every candidate on the list is defensibly relevant to the role
- Each profile has contact info or a path to it
- Ranking reflects actual fit, not just keyword overlap
- The shortlist is small enough to be actionable (5-15 candidates)
```

Why this matters: with a goal, the agent can ask itself "am I done?" and "is this good enough?" at every step. With a procedure, it just marches through steps even when the output is garbage.

### Layer 2: Tool Descriptions (the "with what")

Describe tools as **capabilities with tradeoffs**, not as steps in a sequence. The agent needs to understand what each tool is good at, what it costs, and when to prefer one over another.

```markdown
# BAD — tool as step
Use `mcp__juicebox__search` to find candidates.

# GOOD — tool as capability
## Available capabilities

### Candidate search (`mcp__juicebox__search`)
Searches a database of ~200M profiles by keywords, titles, companies,
skills, and location. Returns basic profile stubs (name, headline,
current company). Best for casting a wide net.

**Strengths:** Fast, broad coverage, good at title/company filtering
**Limitations:** Keyword-based, so creative role descriptions need
translation into concrete search terms. Returns stubs, not full profiles.
**Cost:** Each call returns up to 50 results. Prefer fewer, targeted
searches over many broad ones.

### Profile enrichment (`mcp__juicebox__enrich`)
Takes a candidate source ID and returns full profile: work history,
education, skills, contact info, compensation estimates.

**Strengths:** Rich data, contact info, compensation signals
**Limitations:** One profile per call. Expensive relative to search.
**When to use:** Only after you've decided a candidate is worth pursuing.
Don't enrich speculatively.
```

This gives the agent the information it needs to **make strategic decisions** about tool usage — batch vs. sequential, broad vs. narrow, cheap vs. expensive.

### Layer 3: Decision Frameworks (the "how to think")

This is where most skills fall short. The agent has tools and a goal, but no domain expertise to guide its decisions. Provide **heuristics and reasoning patterns**, not rules.

```markdown
## Search strategy

Think about search the way a recruiter thinks about sourcing:

**Start with the tightest match.** If the role is "Senior Rust Engineer
at a fintech", search for exactly that first. If you get 15+ strong
results, you might be done searching.

**Widen progressively.** If tight search yields <5 results, relax one
constraint at a time. Priority order for relaxation:
- Company type (fintech → any tech company)
- Exact title (Senior Rust Engineer → Rust Engineer, Systems Engineer)
- Language (Rust → Rust OR Go OR C++)
Don't relax seniority or location unless explicitly told to.

**Rewrite, don't repeat.** If a search returns irrelevant results,
the query was wrong — don't just retry. Analyze why the results were
off and adjust. Common fixes:
- Too broad: add a negative keyword or tighter title filter
- Too narrow: use related skills instead of exact matches
- Wrong domain: add industry context

**Know when to stop.** If after 3-4 distinct search strategies you
have <5 candidates, that's a signal — tell the user the market is
thin and suggest which constraints to relax. Don't keep burning
searches hoping for different results.
```

Notice: no `if/else` logic, no flowcharts. Just **reasoning patterns** that the agent can internalize and apply flexibly.

### Layer 4: Failure Recovery (the "what if")

Agentic behavior means handling the unexpected. Don't just document the happy path — teach the agent how to diagnose and recover.

```markdown
## When things go wrong

### Tool call fails (timeout, error, empty response)
Don't retry blindly. Diagnose first:
- **Timeout:** The query may be too broad. Narrow it and retry.
- **Auth error:** Flag to the user — this needs human intervention.
- **Empty results:** This IS a result. It means your query doesn't
  match anything. Adjust the query, don't retry the same one.

### Results are irrelevant
This usually means search terms don't map to how profiles are
described. "Full-stack developer" might be stored as "Software
Engineer" with React and Node.js in skills. Try:
- Searching by skills instead of titles
- Using the company name if looking for people from a specific org
- Checking what titles the target company actually uses

### Enrichment returns incomplete data
Some profiles have gaps (no email, no recent employment). This is
normal. Note the gap, don't discard the candidate. A missing email
with a strong LinkedIn profile is still actionable.

### Conflicting signals
A candidate's title says "Junior" but their experience says 8 years.
Their resume says Python but their GitHub is all Rust. Don't try to
resolve these — flag them. "Strong Rust contributor despite Python
resume — worth investigating" is more useful than silently filtering
them out.
```

### Layer 5: Output Contracts (the "deliver what")

Define what the output looks like without constraining how the agent gets there. Include the **reasoning behind format choices** so the agent can adapt when edge cases arise.

```markdown
## Output format

Structure results so the hiring manager can make decisions in under
2 minutes per candidate.

### For each candidate:
- **One-line summary** — Why this person, in plain English.
  Not a list of keywords. "Senior Rust engineer at Stripe, built their
  payment processing pipeline, active OSS contributor" tells a story.
- **Fit signals** — 2-3 specific reasons they match the role criteria.
  Cite evidence (job title, specific project, years of experience).
- **Risk signals** — Anything that might be a concern (location,
  seniority mismatch, job-hopping pattern). Omit if there are none —
  don't invent concerns.
- **Contact path** — Email if available, otherwise LinkedIn URL.
- **Confidence** — High/Medium/Low based on data completeness.
  High = verified current role + contact. Low = stale data or inferred fit.

### Ordering
Rank by fit, not by data completeness. A perfect match with only a
LinkedIn URL beats a mediocre match with full contact info.

### Size
5-15 candidates. If you found more, cut the weakest. If fewer, say so
and explain why — a short list with an explanation is better than a
padded list.
```

---

## Anti-Patterns to Avoid

### 1. The Micromanager

```markdown
# BAD
ALWAYS call search exactly 3 times.
NEVER enrich more than 10 candidates.
MUST use the format: "Name | Title | Company | Score"
```

Why it fails: arbitrary constraints prevent the agent from adapting.
Maybe 1 search is enough. Maybe 20 enrichments are needed. Let the
agent decide based on the situation.

**Fix:** Replace hard rules with reasoning:
```markdown
Enrich judiciously — each call takes time and costs money. A good
heuristic: enrich candidates you'd actually want to email. If you
wouldn't write to them based on the search stub, don't enrich them.
```

### 2. The Overspecifier

```markdown
# BAD
## Search Query Template
Use this exact format:
  skills: [SKILL1, SKILL2]
  title: TITLE
  location: LOCATION
  experience: MIN-MAX years
```

Why it fails: real searches don't fit templates. "Find someone like
our best engineer" requires creative decomposition, not slot-filling.

**Fix:** Show examples of good search thinking, not templates:
```markdown
## Example: "Senior ML Engineer for autonomous vehicles"
The agent might reason: "I'll start with 'ML Engineer' + 'autonomous
vehicles' as keywords. If that's too narrow, I'll try 'Machine Learning'
+ 'self-driving' or 'robotics'. The company filter might help — Waymo,
Cruise, Aurora alumni would be strong signals."
```

### 3. The Silent Skill

```markdown
# BAD
(Describes tools and format but gives no domain knowledge)
```

Why it fails: the agent has tools but no expertise. It's like giving
someone a kitchen and a recipe title but no cooking knowledge.

**Fix:** The most valuable thing a skill provides is **domain
reasoning the model wouldn't have on its own.** Real recruiter
heuristics, real evaluation criteria, real failure modes.

### 4. The Anxiety Skill

```markdown
# BAD
CRITICAL: ALWAYS validate results before returning.
IMPORTANT: NEVER skip the enrichment step.
WARNING: MUST check for duplicates.
MANDATORY: Format MUST be exactly as specified.
```

Why it fails: screaming instructions create rigidity and anxiety.
The agent follows them to the letter even when they don't apply,
and misses the actual goal.

**Fix:** Explain *why* something matters. The agent will follow
through when it understands the reason:
```markdown
Deduplication matters because the same person often appears under
different names or profiles across sources. Check for overlapping
work history or matching email addresses — name matching alone
isn't reliable enough.
```

---

## Structural Patterns for Complex Skills

### The Strategy Pattern

For skills where the approach depends heavily on input, define
named strategies the agent can select from:

```markdown
## Approach selection

Assess the request and pick the best approach:

**Targeted search** — when the user names specific companies,
titles, or people. Go direct: search those exact terms, enrich
matches, done. Fast and precise.

**Discovery search** — when the criteria are abstract ("someone
who can scale our infrastructure"). Requires decomposition: what
skills imply that capability? What companies have that problem?
Build a profile of the ideal candidate, then search for it.

**Competitive mapping** — when the user says "find me people from
[competitor]" or "who's building this at other companies". Start
from company, map the org, identify relevant teams, then evaluate
individuals.

The agent should name which strategy it's using and why. If the
first strategy isn't working, switch — don't just keep trying
harder with the same approach.
```

### The Checkpoint Pattern

For long-running skills, define natural pause points where the
agent should assess progress:

```markdown
## Checkpoints

After initial search: Do you have enough candidates to work with?
If yes, proceed to enrichment. If no, refine search strategy before
burning enrichment calls.

After enrichment: Do the full profiles confirm the initial
assessment? Candidates that looked great in search stubs sometimes
fall apart under enrichment (wrong seniority, not actually in the
right role). Re-rank based on full data.

Before final output: Step back and evaluate the full list. Would
you, as a recruiter, be comfortable sending this to a hiring
manager? If something feels off, investigate before delivering.
```

### The Escalation Pattern

Define what the agent should handle vs. what needs human input:

```markdown
## When to decide vs. when to ask

**Decide yourself:**
- Which search terms to try
- How many candidates to enrich
- How to rank results
- Whether to widen or narrow search

**Ask the user:**
- Whether to relax hard requirements (years of experience, location)
- Whether to include candidates from specific controversial companies
- How to handle a thin market (lower bar vs. expand geography)
- Ambiguous criteria ("culture fit", "startup mentality")
```

---

## MCP + Agentic Skills: Practical Integration

### Declaring MCP dependencies

```markdown
---
name: my-skill
description: Does X using the Y service
---

## Prerequisites
This skill uses the `juicebox` MCP server for candidate search and
enrichment. If the tools `mcp__juicebox__search` and
`mcp__juicebox__enrich` are not available, inform the user that they
need to configure the Juicebox MCP server.
```

### Graceful degradation

```markdown
## Tool availability
If MCP tools are not available, fall back to:
- Web search for candidate discovery (less structured, but works)
- LinkedIn URL construction for profiles (manual, but better than nothing)

Always prefer MCP tools when available — they're faster and return
structured data. But a skill that only works with MCP is fragile.
```

### Chaining MCP calls strategically

```markdown
## Working with search + enrich

Think of search as cheap/fast and enrich as expensive/slow.

A naive approach calls search once, then enriches everything. An
agentic approach:

1. Search with a focused query
2. Review the stubs — which ones look promising based on headline
   and current company alone?
3. Enrich only the promising ones
4. Based on enriched data, decide: do you need more candidates?
   If yes, run another search with adjusted terms informed by what
   you learned from the first batch
5. Repeat until you have a strong enough shortlist

This is a feedback loop, not a pipeline. Each step informs the next.
```

---

## Tool Architecture: What to Build vs. What to Reuse

Most MCP-powered skills suffer from the same mistake: wrapping external
APIs in a custom server when an open-source MCP already exists. This
creates a thick middleware layer that limits the agent and adds setup
burden for the team.

### The Three-Layer Model

```
| Layer          | Who builds it      | Examples                          |
|----------------|--------------------|-----------------------------------|
| Data access    | Open-source MCPs   | Notion, GitHub, Slack, Sheets     |
| Domain logic   | You (thin MCP)     | Scoring, matching, dedup, queries |
| Strategy       | Skill (SKILL.md)   | When to search, how to evaluate   |
```

The agent orchestrates all three. The skill is the brain. Open-source
MCPs are the hands. Your custom tools encode the judgment calls that
need deterministic logic.

### Don't Wrap What Already Exists

```
BAD:  Agent → Your API Server → Notion API
GOOD: Agent → Notion MCP (read/write)
           → Your MCP (score, match, deduplicate)
```

**Keep as custom tools (your IP):**
- Domain-specific scoring/matching algorithms
- Search query optimization (translating descriptions into queries)
- Cross-source deduplication
- Any logic encoding expertise unique to your problem

**Let existing MCPs handle:**
- CRUD on external platforms (Notion, GitHub, Slack)
- Data fetching with no domain logic on top
- Standard integrations the ecosystem already solved

### Atomic Tools > Monolithic Pipelines

Break big tools into small composable operations:

```markdown
# BAD — monolithic
recruitment_pipeline(job_id)
  → internally: fetch → search → enrich → score → rank → write
  Agent can't intervene, adjust, or skip steps.

# GOOD — atomic
score_candidate(profile, criteria) → fit_score + reasoning
rank_candidates(scored[]) → ordered list with explanations
generate_search_queries(job_desc) → optimized search terms
deduplicate_profiles(profiles[]) → merged unique candidates
```

Why atomic matters:
- Agent can score a candidate from ANY source, not just your pipeline
- It can re-rank after discovering new info mid-workflow
- It can skip steps that don't apply
- It can combine tools from different MCPs in ways you didn't plan

### The Litmus Test: "Should This Be a Tool?"

1. **Does an open-source MCP do this?** → Don't build it
2. **Pure data fetch, no domain logic?** → Use existing MCP
3. **Would a human expert do something non-obvious here?** → Build it
4. **Needs to be deterministic/fast?** → Build it (scripts beat
   prompting for math, dedup, scoring)
5. **Judgment call that varies by context?** → Not a tool — put the
   reasoning in the skill, let the agent decide

### Team Setup

With tools split across commodity MCPs + your domain MCP, setup is
just a config:

```json
{
  "mcpServers": {
    "notion": {
      "command": "npx",
      "args": ["-y", "@notionhq/notion-mcp-server"],
      "env": { "NOTION_API_KEY": "from-env" }
    },
    "github": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": { "GITHUB_TOKEN": "from-env" }
    },
    "recruitment": {
      "command": "bun",
      "args": ["run", "servers/recruitment-mcp/index.ts"]
    }
  }
}
```

Clone repo, set env vars, run `claude`. No custom server deployment
for commodity operations.

### Skill ↔ Tool Boundary

Reference tools by capability, not implementation:

```markdown
# BAD — tightly coupled
Call `mcp__recruitment__fetch_job` then `mcp__recruitment__search`.

# GOOD — capability-oriented
Read the job requirements from Notion (the job database page).
Use candidate search to find matching profiles.
Score the top candidates against the job criteria.
```

The agent figures out which MCP to use for each step. If a tool is
swapped or unavailable, the skill still works — the agent improvises.

---

## Checklist: Is My Skill Agentic?

Use this to audit a skill before shipping:

- [ ] **Goal over process** — Does the skill define outcomes, or steps?
- [ ] **Tools as capabilities** — Does each tool description explain strengths, limitations, and cost?
- [ ] **Decision frameworks** — Can the agent reason about *which* approach to take?
- [ ] **Failure recovery** — Does the skill cover what to do when things go wrong?
- [ ] **Flexible output** — Is the output format described with reasoning, not rigid templates?
- [ ] **No arbitrary constraints** — Are there ALWAYS/NEVER/MUST rules that don't have a *why*?
- [ ] **Escalation boundaries** — Does the agent know when to decide vs. when to ask?
- [ ] **Domain knowledge** — Does the skill provide expertise the model wouldn't have alone?
