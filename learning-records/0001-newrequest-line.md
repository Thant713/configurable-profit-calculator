# 0001 — First session: the http.NewRequest line

**Date:** 2026-08-21
**Target:** explain `req, err := http.NewRequest(...)` + full fetchItemPrice body

## Learned (can explain back)
- `NewRequest` builds request struct in RAM only; returns `(handle-pointer, error)`
- Pointer = address/handle; `req` passed to `Do`
- Two error checks = different failure classes: build-time text problems vs transport/world problems
- Misspelled domain → Do's error (parses fine); syntax garbage → NewRequest's error
- Status check = third failure class: delivered but rejected; `StatusCode` int vs `Status` string
- `%d`=int, `%s`=string format slots; `fmt.Errorf` builds errors
- Decode = unmarshalling (user connected this themselves)

## Fuzzy / not yet explained back
- Map lookup with `ok` idiom (`p, ok := r.Data[strconv.Itoa(id)]`)
- Why `&r` for Decode (address-of so decoder can fill)
- `strconv.Itoa` role in lookup key (IDs are ints, map keys strings)

## Next session (ZPD)
- Start: map lookup + ok idiom quiz
- Then: caller side — main.go calling fetchItemPrice for all 4 items
- Then: profit calc (item.go)

## Notes
- User explains first, agent guides via questions — keep this reversed flow
- Voice transcription garbles; re-confirm terms when odd words appear
